package middleware

import (
	cfg "gm-exam/config"
	auth "gm-exam/src/authenticator"
	"strings"

	"github.com/kataras/iris/v12"
)

func ValidateToken(ctx iris.Context) {
	authHeader := ctx.GetHeader("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" || authHeader == "" {
		ctx.StopWithProblem(iris.StatusTokenRequired, iris.NewProblem().Title(cfg.InvalidTokenType))
		return
	}

	valid, session, err := auth.ValidateToken(authHeaderParts[1])
	if err != nil {
		ctx.StopWithProblem(iris.StatusInvalidToken, iris.NewProblem().Title(err.Error()))
		return
	}

	if !valid {
		ctx.StopWithProblem(iris.StatusInvalidToken, iris.NewProblem().Title(cfg.InvalidToken))
		return
	}

	ctx.Values().Set("session", session)
	ctx.Next()
}
