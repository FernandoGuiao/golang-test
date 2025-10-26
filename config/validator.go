package config

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptbr_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var trans ut.Translator

func SetupValidator() {
	ptBR := pt_BR.New()
	uni := ut.New(ptBR, ptBR)

	ttrans, _ := uni.GetTranslator("pt_BR")
	trans = ttrans

	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		_ = ptbr_translations.RegisterDefaultTranslations(validate, ttrans)
	}
}

func GetTranslator() ut.Translator {
	return trans
}
