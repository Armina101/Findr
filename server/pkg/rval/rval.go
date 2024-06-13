package rval

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translate "github.com/go-playground/validator/v10/translations/en"
)

var timeOut = 10 * time.Second

var errs []error

// FieldValidator create a contextual validation of all the struct field and their struct
// tags and proper configuration of the validation error
func FieldValidator(s interface{}) []error {
	vald := validator.New()
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeOut)
	defer cancelCtx()

	// setting english translator
	eng := en.New()
	translate := ut.New(eng, eng)
	translator, ok := translate.GetTranslator("en")

	if !ok {
		fmt.Println("unable to set english translator")
		return []error{fmt.Errorf("cannot process validation of input")}
	}

	err := en_translate.RegisterDefaultTranslations(vald, translator)
	if err != nil {
		return []error{fmt.Errorf("cannot register default translation")}
	}

	err = vald.StructCtx(ctx, s)
	var fieldErrors validator.ValidationErrors
	errors.As(err, &fieldErrors)

	for _, v := range fieldErrors {
		err := fmt.Errorf(v.Translate(translator))
		errs = append(errs, err)
	}

	// fmt.Println("FieldValidatorError: error invalid input validation")
	return errs
}

// ValidateMail: Validate user email input using regex expression
func ValidateMail(email string) bool {
	regMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	ok := regMail.MatchString(email)
	return ok
}