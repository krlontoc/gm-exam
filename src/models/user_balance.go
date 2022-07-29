package models

import (
	"errors"
	cfg "gm-exam/config"
	"time"

	"gorm.io/gorm"
)

type UserBalance struct {
	tx        *gorm.DB
	ID        uint      `json:"id" gomr:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID  uint    `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (this *UserBalance) dbCheck() error {
	if cfg.DB == nil {
		return errors.New(cfg.DBError)
	}

	if this.tx == nil {
		this.tx = cfg.DB
	}

	return nil
}

func (this *UserBalance) InTx(tx *gorm.DB) {
	this.tx = tx
}

func (this *UserBalance) Get(query string, args []interface{}) error {
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
