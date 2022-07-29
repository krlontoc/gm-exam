package controllers

import (
	"fmt"

	vldtr "github.com/go-playground/validator/v10"
)

func ReadValidatorMessage(errs vldtr.ValidationErrors) []string {
	errMsg := []string{}
	for _, err := range errs {
		errString := fmt.Sprintf("%s is %s.", err.Field(), err.Tag())
		errMsg = append(errMsg, errString)
	}
	return errMsg
}
