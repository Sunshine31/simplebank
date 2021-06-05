package api

import (
	"github.com/Sunshine31/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"CAD": true,
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
