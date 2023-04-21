package main

import (
	cfg "gm-exam/config"
	intl "gm-exam/initial"
	auth "gm-exam/src/authenticator"
	mdlw "gm-exam/src/middlewares"
	"os"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	tnx "gm-exam/src/controllers/transaction"
	usr "gm-exam/src/controllers/user"

	_ "gm-exam/docs"
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

	// Configure the swagger UI page.
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("http://localhost:1007/swagger/doc.json"), // The url pointing to docs definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)

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

// @title GM Technical Exam
// @description This is my GoMedia technical exam API docs.

// @host 35.247.166.232
// @BasePath /api/v1
func main() {
	app := initApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "1007"
	}
	app.Listen(":" + port)
}
