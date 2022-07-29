package transaction

import (
	"fmt"
	"strings"

	vldtr "github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	cfg "gm-exam/config"
	ctrl "gm-exam/src/controllers"
	mdls "gm-exam/src/models"
)

func Deposit(ctx iris.Context) {
	form := &mdls.DepositForm{}
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

	ss := ctx.Values().Get("session").(mdls.User)
	form.UserID = ss.ID
	tnx, err := mdls.Transaction{}.Deposit(*form)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(cfg.InternalError).Detail(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": tnx})
}

func Send(ctx iris.Context) {
	ss := ctx.Values().Get("session").(mdls.User)
	form := &mdls.SendForm{}
	if err := ctx.ReadJSON(form); err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	if form.To == ss.EmailAddress {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidRecipient).Detail("Can not use [self] as recipient."))
		return
	}

	validator := vldtr.New()
	err := validator.Struct(form)
	if err != nil {
		errMsg := ctrl.ReadValidatorMessage(err.(vldtr.ValidationErrors))
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.RequiredField).Detail(strings.Join(errMsg, " ")))
		return
	}

	cBalance, err := ss.GetBalance(0)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(cfg.InternalError).Detail(err.Error()))
		return
	}
	if cBalance < form.Amount {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InsufficientBalance).Detail(fmt.Sprintf("Current balance is %v.", cBalance)))
		return
	}

	form.From = ss.ID
	tnx, err := mdls.Transaction{}.Send(*form)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(cfg.InternalError).Detail(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": tnx})
}

func MultiSend(ctx iris.Context) {
	form := &mdls.MultiSendForm{}
	if err := ctx.ReadJSON(form); err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	if len(form.Sends) <= 0 {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.RequiredField))
		return
	}

	validator := vldtr.New()
	err := validator.Struct(form)
	if err != nil {
		errMsg := ctrl.ReadValidatorMessage(err.(vldtr.ValidationErrors))
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.RequiredField).Detail(strings.Join(errMsg, " ")))
		return
	}

	ss := ctx.Values().Get("session").(mdls.User)
	cBalance, err := ss.GetBalance(0)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(cfg.InternalError).Detail(err.Error()))
		return
	}

	tAmount := 0.0
	for _, sF := range form.Sends {
		tAmount += sF.Amount
	}
	if cBalance < tAmount {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InsufficientBalance).Detail(fmt.Sprintf("Current balance is %v, and total amount needed is %v.", cBalance, tAmount)))
		return
	}

	form.From = ss.ID
	tnxs, err := mdls.Transaction{}.MultiSend(*form)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().Title(cfg.InternalError).Detail(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": tnxs})
}
