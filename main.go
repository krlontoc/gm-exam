package main

import (
	cfg "gm-exam/config"
	intl "gm-exam/initial"
	auth "gm-exam/src/authenticator"
	mdlw "gm-exam/src/middlewares"

	"github.com/kataras/iris/v12"

	tnx "gm-exam/src/controllers/transaction"
	usr "gm-exam/src/controllers/user"
)

func init() {
	cfg.Init()
	intl.InitialSetup()
}

func initApp() *iris.Application {
	app := iris.New()

	// Root
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": iris.StatusOK, "message": "Hi! This is GoMedia Technical Exam"})
	})

	// Public Endpoints
	ath := app.Party("/auth/v1")
	{
		ath.Handle("POST", "/login", auth.Login)
	}

	// Authenticated Endpoints
	api := app.Party("/api/v1", mdlw.ValidateToken)
	{
		api.Handle("GET", "/user/{ID:uint}", usr.Get)
		api.Handle("GET", "/user/balance", usr.GetBalance)
		api.Handle("GET", "/user/transactions", usr.GetTransactionList)
		api.Handle("POST", "/transaction/deposit", tnx.Deposit)
		api.Handle("POST", "/transaction/send", tnx.Send)
		api.Handle("POST", "/transaction/multi-send", tnx.MultiSend)
	}

	// Ease of Use
	app.Get("/users", usr.List)

	return app
}

func main() {
	app := initApp()
	app.Listen(":1007")
}
