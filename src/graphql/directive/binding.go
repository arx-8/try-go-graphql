package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
}

// Binding ... @binding custom directive
// See https://david-yappeter.medium.com/gqlgen-custom-data-validation-part-1-7de8ef92de4c
func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}
	fieldName := *graphql.GetPathContext(ctx).Field

	err = validate.Var(val, constraint)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		transErr := fmt.Errorf("%s%+v", fieldName, validationErrors[0].Translate(trans))
		return val, transErr
	}

	return val, nil
}
