package authenticator

import (
	ctrl "gm-exam/src/controllers"
	"strings"

	vldtr "github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	cfg "gm-exam/config"
	mdls "gm-exam/src/models"
)

type LogInForm struct {
	EmailAddress string `json:"email_address" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

func Login(ctx iris.Context) {
	form := &LogInForm{}
	if err := ctx.ReadJSON(form); err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	validator := vldtr.New()
	err := validator.Struct(form)
	if err != nil {
		errMsg := ctrl.ReadValidatorMessage(err.(vldtr.ValidationErrors))
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.RequiredField).Detail(strings.Join(errMsg, " ")))
		return
	}

	user, err := mdls.User{}.Login(form.EmailAddress, form.Password)
	if err != nil || user == nil || user.ID == 0 {
		ctx.StopWithProblem(iris.StatusOK, iris.NewProblem().Title(cfg.InvalidLogin))
		return
	}

	token, err := GenerateToken(*user)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(err.Error()))
		return
	}
	ctx.JSON(iris.Map{"code": iris.StatusOK, "data": token})
}
