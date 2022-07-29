package main

import (
	auth "gm-exam/src/authenticator"
	mdls "gm-exam/src/models"
	"log"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

// Test Route Details
type TRD struct {
	Path      string
	Method    string
	Body      map[string]interface{}
	EStatus   int      // expected status
	EResponse iris.Map // expected response
}

func TestInitApp(t *testing.T) {
	// generate token for middleware test
	usr := mdls.User{}
	usr.Get("id = ?", []interface{}{1})
	token, err := auth.GenerateToken(usr)
	if err != nil {
		log.Panic("Unable to generate token for testing")
	}

	a := initApp()
	e := httptest.New(t, a)

	resp := map[string]interface{}{
		"status":  iris.StatusOK,
		"message": "Hi! This is GoMedia Technical Exam",
	}
	e.GET("/").Expect().Status(iris.StatusOK).JSON().Equal(resp)
	e.GET("/users").Expect().Status(iris.StatusOK).JSON().Object().Keys().Contains("data", "status")

	body := map[string]interface{}{
		"email_address": "kr@email.com",
		"password":      "Pass123!",
	}
	e.POST("/auth/v1/login").WithJSON(body).Expect().Status(iris.StatusOK)

	resp = map[string]interface{}{
		"status": iris.StatusOK,
		"data": map[string]interface{}{
			"id":            1,
			"full_name":     "Kurt Russel",
			"email_address": "kr@email.com",
		},
	}
	e.GET("/api/v1/user/1").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		Expect().Status(iris.StatusOK).JSON().Object().ContainsMap(resp)

	e.GET("/api/v1/user/balance").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		Expect().Status(iris.StatusOK).JSON().Object().Keys().Contains("data", "status")

	e.GET("/api/v1/user/transactions").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		Expect().Status(iris.StatusOK).JSON().Object().Keys().Contains("data", "status")

	body = map[string]interface{}{"amount": 10000}
	resp = map[string]interface{}{
		"status": iris.StatusOK,
		"data": map[string]interface{}{
			"to":      1,
			"amount":  10000,
			"message": "Deposit Transaction",
		},
	}
	e.POST("/api/v1/transaction/deposit").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		WithJSON(body).Expect().Status(iris.StatusOK).JSON().Object().ContainsMap(resp)

	body = map[string]interface{}{
		"to":      "k@email.com",
		"amount":  5000,
		"message": "Rent Payment",
	}
	resp = map[string]interface{}{
		"status": iris.StatusOK,
		"data": map[string]interface{}{
			"from":    1,
			"to":      "k@email.com",
			"amount":  5000,
			"message": "Rent Payment",
		},
	}
	e.POST("/api/v1/transaction/send").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		WithJSON(body).Expect().Status(iris.StatusOK).JSON().Object().ContainsMap(resp)

	body = map[string]interface{}{
		"sends": []map[string]interface{}{
			{
				"to":      "kl@email.com",
				"amount":  2500,
				"message": "Allowance",
			},
			{
				"to":      "rl@email.com",
				"amount":  2500,
				"message": "Allowance",
			},
		},
	}
	resp = map[string]interface{}{
		"status": iris.StatusOK,
		"data": []map[string]interface{}{
			{
				"send": map[string]interface{}{
					"to":      "kl@email.com",
					"amount":  2500,
					"message": "Allowance",
				},
				"status": "OK",
				"title":  "Success",
			},
			{
				"send": map[string]interface{}{
					"to":      "rl@email.com",
					"amount":  2500,
					"message": "Allowance",
				},
				"status": "OK",
				"title":  "Success",
			},
		},
	}
	e.POST("/api/v1/transaction/multi-send").WithHeaders(map[string]string{"Authorization": "Bearer " + token}).
		WithJSON(body).Expect().Status(iris.StatusOK).JSON().Object().ContainsMap(resp)
}
