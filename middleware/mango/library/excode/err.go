package excode

import "github.com/x554462/go-exception"

var ValidateError = exception.New("验证错误", exception.RootError)
var RuntimeError = exception.New("服务器内部错误", exception.RootError)
