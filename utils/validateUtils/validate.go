package validateUtils

import (
	"github.com/go-playground/validator/v10"
	"index_Demo/bootstrap/global"
	"reflect"
)

// GetValidateErrMessages 获取验证器错误
func GetValidateErrMessages(st any, err error) map[string]string {
	t := reflect.TypeOf(st)

	//没有错误直接返回
	if err == nil {
		return map[string]string{}
	}

	errors, ok := err.(validator.ValidationErrors)

	//验证器没有错误
	if !ok {
		return map[string]string{}
	}

	//验证器有错误
	results := make(map[string]string)
	for _, e := range errors {
		//logger.SugarLogger.Errorf("e -> %s", e.Tag())
		name, _ := t.Elem().FieldByName(e.Field())
		results[e.Tag()+"_msg"] = name.Tag.Get(e.Tag() + "_msg")
	}
	return results
}

// ReturnValidateMessage 返回验证器的错误
func ReturnValidateMessage(st any, err error) (string, bool) {
	if err == nil {
		//无错误
		return "", false
	}

	//验证器错误
	messages := GetValidateErrMessages(st, err)
	if len(messages) > 0 {
		for _, msg := range messages {
			return msg, true
		}
	}

	global.SugarLogger.Errorf("et -> %s", reflect.TypeOf(err).String())

	//绑定错误
	return err.Error(), true
}
