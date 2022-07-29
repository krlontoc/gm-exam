package models

import (
	"errors"
	"time"

	cfg "gm-exam/config"

	"gorm.io/gorm"
)

type User struct {
	tx        *gorm.DB
	ID        uint      `json:"id" gomr:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	FullName     string `json:"full_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	Password     string `json:"password,omitempty" gorm:"->:false;<-"`
}

func (this *User) dbCheck() error {
	if cfg.DB == nil {
		return errors.New(cfg.DBError)
	}

	if this.tx == nil {
		this.tx = cfg.DB
	}

	return nil
}

func (this *User) InTx(tx *gorm.DB) {
	this.tx = tx
}

func (this *User) Get(query string, args []interface{}) error {
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

func (this User) Find(query string, args []interface{}, qc *cfg.QueryConfig) ([]User, error) {
	if err := this.dbCheck(); err != nil {
		return []User{}, err
	}

	data := []User{}
	if qc != nil && len(qc.Columns) > 0 {
		this.tx = this.tx.Select(qc.Columns)
	}
	this.tx = this.tx.Where(query, args...)
	if qc != nil {
		this.tx = cfg.SetQueryConfig(this.tx, *qc)
	}
	if err := this.tx.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (this User) Login(email, pass string) (*User, error) {
	if err := this.Get("email_address = ? AND password = ?", []interface{}{email, pass}); err != nil {
		return nil, err
	}

	return &this, nil
}

func (this *User) GetBalance(id uint) (float64, error) {
	if err := this.dbCheck(); err != nil {
		return 0, err
	}

	// prioritize targeted query
	if id > 0 {
		this.ID = id
	}

	ub := UserBalance{}
	ub.InTx(this.tx)
	if err := ub.Get("user_id = ?", []interface{}{this.ID}); err != nil {
		return 0.0, err
	}

	return ub.Balance, nil
}

// tnxType: 0: All, 1:Send, 2:Receive
func (this *User) GetTransactions(tnxType int, id uint) ([]Transaction, error) {
	if err := this.dbCheck(); err != nil {
		return []Transaction{}, err
	}

	// prioritize targeted query
	if id > 0 {
		this.ID = id
	}

	tnxQuery := "`from` = ? OR `to` = ?"
	switch tnxType {
	case 1:
		tnxQuery = "`from` = ?"
	case 2:
		tnxQuery = "`to` = ?"
	}

	tnx := Transaction{}
	tnx.InTx(this.tx)

	qCfg := cfg.QueryConfig{
		Order:   "DESC",
		OrderBy: "`created_at`",
	}
	tnxs, err := tnx.Find(
		tnxQuery,
		[]interface{}{this.ID, this.ID},
		&qCfg,
	)
	if err != nil {
		return nil, err
	}

	return tnxs, nil
}

func (this User) UpdateBalance(id uint, amount float64) error {
	if err := this.dbCheck(); err != nil {
		return err
	}

	cBal, err := this.GetBalance(id)
	if err != nil {
		return err
	}

	nBal := cBal + amount
	if err := this.tx.Model(&UserBalance{}).Where("`user_id` = ?", id).Update("balance", nBal).Error; err != nil {
		return err
	}

	return nil
}
