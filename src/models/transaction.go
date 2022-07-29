package models

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"

	cfg "gm-exam/config"
)

type Transaction struct {
	tx        *gorm.DB
	ID        uint      `json:"id" gomr:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	From    uint    `json:"from"`
	To      uint    `json:"to"`
	Amount  float64 `json:"amount"`
	Message string  `json:"message" gorm:"type:text"`
}

type DepositForm struct {
	UserID uint
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type SendForm struct {
	From    uint    `json:"from,omitempty"`
	To      string  `json:"to" validate:"required"`
	Amount  float64 `json:"amount" validate:"required,gt=0"`
	Message string  `json:"message" validate:"required"`
}

type MultiSendForm struct {
	From  uint
	Sends []SendForm `json:"sends" validate:"gte=0,required,dive,required"`
}

type MultiSendResponse struct {
	Send    SendForm `json:"send,omitempty"`
	Status  string   `json:"status,omitempty"`
	Title   string   `json:"title,omitempty"`
	Details string   `json:"details,omitempty"`
}

func (this *Transaction) dbCheck() error {
	if cfg.DB == nil {
		return errors.New(cfg.DBError)
	}

	if this.tx == nil {
		this.tx = cfg.DB
	}

	return nil
}

func (this *Transaction) InTx(tx *gorm.DB) {
	this.tx = tx
}

func (this *Transaction) Get(query string, args []interface{}) error {
	if err := this.dbCheck(); err != nil {
		return err
	}

	var err error = nil
	if this.ID > 0 {
		err = this.tx.First(this, this.ID).Error
	}

	if query != "" {
		err = this.tx.Where(query, args...).First(this).Error
	}

	return err
}

func (this Transaction) Find(query string, args []interface{}, qc *cfg.QueryConfig) ([]Transaction, error) {
	if err := this.dbCheck(); err != nil {
		return []Transaction{}, err
	}

	data := []Transaction{}
	if qc != nil && len(qc.Columns) > 0 {
		this.tx = this.tx.Select(qc.Columns)
	}
	this.tx = this.tx.Where(query, args...)
	if qc != nil {
		this.tx = cfg.SetQueryConfig(this.tx, *qc)
	}
	if err := this.tx.Find(&data).Error; err != nil {
		return data, cfg.DB.Error
	}

	return data, nil
}

func (this Transaction) Deposit(form DepositForm) (Transaction, error) {
	if err := this.dbCheck(); err != nil {
		return Transaction{}, err
	}

	// set values
	this.To = form.UserID
	this.Amount = form.Amount
	this.Message = cfg.DepositTransaction

	// begin DB transaction
	tx := this.tx.Begin()

	// creaste transaction record
	if err := tx.Create(&this).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	usr := User{}
	usr.InTx(tx)

	// update user balanace
	if err := usr.UpdateBalance(form.UserID, form.Amount); err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// commit DB transaction
	if err := tx.Commit().Error; err != nil {
		return Transaction{}, err
	}

	return this, nil
}

func (this Transaction) Send(form SendForm) (SendForm, error) {
	if err := this.dbCheck(); err != nil {
		return form, err
	}

	// set values
	this.ID = 0
	this.From = form.From
	this.Amount = form.Amount
	this.Message = form.Message

	to := User{}
	if err := to.Get("`email_address` = ?", []interface{}{form.To}); err != nil || to.ID == 0 {
		return form, errors.New(cfg.InvalidRecipient)
	}
	if to.ID == form.From {
		return form, errors.New("Can not use [self] as recipient.")
	}
	this.To = to.ID

	// begin DB transaction
	tx := this.tx.Begin()

	// creaste transaction record
	if err := tx.Create(&this).Error; err != nil {
		tx.Rollback()
		return form, err
	}

	usr := User{}
	usr.InTx(tx)

	// update sender's balanace
	if err := usr.UpdateBalance(form.From, (form.Amount * -1)); err != nil {
		tx.Rollback()
		return form, err
	}

	// update receipient's balanace
	if err := usr.UpdateBalance(to.ID, form.Amount); err != nil {
		tx.Rollback()
		return form, err
	}

	// commit DB transaction
	if err := tx.Commit().Error; err != nil {
		return form, err
	}

	return form, nil
}

func (this Transaction) MultiSend(form MultiSendForm) ([]MultiSendResponse, error) {
	if err := this.dbCheck(); err != nil {
		return nil, err
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(form.Sends))
	msResp := make([]MultiSendResponse, len(form.Sends))

	// to ensure that the waitgroup will only do 10 concurrent processes
	guard := make(chan struct{}, 10)
	for i, sF := range form.Sends {
		guard <- struct{}{}
		go func(i int, sF SendForm) {
			defer func() {
				<-guard
			}()
			defer wg.Done()

			sF.From = form.From
			rsF, err := (Transaction{}).Send(sF)
			rsF.From = 0
			sResponse := MultiSendResponse{
				Send:   rsF,
				Status: "OK",
				Title:  cfg.Success,
			}
			if err != nil {
				sResponse.Status = "ERROR"
				sResponse.Title = err.Error()
				if err.Error() == "Can not use [self] as recipient." {
					sResponse.Title = cfg.InvalidRecipient
					sResponse.Details = "Can not use [self] as recipient."
				}

			}
			msResp[i] = sResponse

		}(i, sF)
	}

	wg.Wait()

	return msResp, nil
}
