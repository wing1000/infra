package base

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"imooc.com/resk/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	Check(validate)
	return validate
}

func Transtate() ut.Translator {
	Check(translator)
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error("Not found translator: zh")
	}

}

func ValidateStruct(s interface{}) (err error) {
	err = Validate().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			log.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				log.Error(e.Translate(Transtate()))
			}
		}
		return err
	}
	return nil
}
