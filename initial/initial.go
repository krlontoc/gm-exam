package initial

import (
	cfg "gm-exam/config"
	mdls "gm-exam/src/models"
	"time"
)

func InitialSetup() {
	// app was initialized already
	if cfg.DB.Migrator().HasTable(&mdls.User{}) {
		return
	}

	InitTables()
	SeedData()
}

func InitTables() {
	cfg.DB.AutoMigrate(
		&mdls.User{}, &mdls.UserBalance{},
		&mdls.Transaction{},
	)
}

func SeedData() {
	locTZ := cfg.GetLocTZ()

	// Users
	users := []mdls.User{
		{
			FullName:     "Kurt Russel",
			EmailAddress: "kr@email.com",
			Password:     "Pass123!",
			CreatedAt:    time.Now().In(locTZ),
			UpdatedAt:    time.Now().In(locTZ),
		},
		{
			FullName:     "Kurt",
			EmailAddress: "k@email.com",
			Password:     "Pass123!",
			CreatedAt:    time.Now().In(locTZ),
			UpdatedAt:    time.Now().In(locTZ),
		},
		{
			FullName:     "Russel",
			EmailAddress: "r@email.com",
			Password:     "Pass123!",
			CreatedAt:    time.Now().In(locTZ),
			UpdatedAt:    time.Now().In(locTZ),
		},
		{
			FullName:     "Kurt Lontoc",
			EmailAddress: "kl@email.com",
			Password:     "Pass123!",
			CreatedAt:    time.Now().In(locTZ),
			UpdatedAt:    time.Now().In(locTZ),
		},
		{
			FullName:     "Russel Lontoc",
			EmailAddress: "rl@email.com",
			Password:     "Pass123!",
			CreatedAt:    time.Now().In(locTZ),
			UpdatedAt:    time.Now().In(locTZ),
		},
	}
	cfg.DB.Create(&users)

	// User Balance
	userBalances := []mdls.UserBalance{}
	for _, user := range users {
		userBalances = append(userBalances, mdls.UserBalance{
			UserID:    user.ID,
			Balance:   15000.0,
			CreatedAt: time.Now().In(locTZ),
			UpdatedAt: time.Now().In(locTZ),
		})
	}
	cfg.DB.Create(&userBalances)
}
