package user

import (
	"strconv"

	"github.com/kataras/iris/v12"

	cfg "gm-exam/config"
	mdls "gm-exam/src/models"
)

func List(ctx iris.Context) {
	search := ctx.Params().Get("search")

	users, err := (mdls.User{}).Find(
		"full_name LIKE ? OR email_address LIKE ?",
		[]interface{}{"%" + search + "%", "%" + search + "%"},
		&cfg.QueryConfig{
			Columns: []string{"full_name", "email_address"},
		},
	)

	data := []map[string]interface{}{}
	for _, u := range users {
		data = append(data, map[string]interface{}{
			"full_name":     u.FullName,
			"email_address": u.EmailAddress,
		})
	}

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": data})
}

func Get(ctx iris.Context) {
	param := ctx.Params().Get("ID")
	var id uint
	if tmp, err := strconv.Atoi(param); err == nil && tmp > 0 {
		id = uint(tmp)
	} else {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(cfg.InvalidPayload))
		return
	}

	user := mdls.User{}
	user.ID = id
	if err := user.Get("", nil); err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": user})
}

func GetBalance(ctx iris.Context) {
	ss := ctx.Values().Get("session").(mdls.User)
	balance, err := ss.GetBalance(0)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": balance})
}

func GetTransactionList(ctx iris.Context) {
	ss := ctx.Values().Get("session").(mdls.User)
	balance, err := ss.GetTransactions(0, 0)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title(err.Error()))
		return
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": balance})
}
