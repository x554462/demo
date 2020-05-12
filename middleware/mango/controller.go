package mango

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/x554462/demo/middleware/mango/library/excode"
	"github.com/x554462/demo/middleware/mango/library/logging"
	"github.com/x554462/demo/middleware/mango/library/util"
	"github.com/x554462/demo/middleware/mango/validator"
	"github.com/x554462/go-exception"
	"github.com/x554462/sorm"
	"net/http"
	"strings"
)

const DefaultKey = "middleware/mango"

type Controller struct {
	ginCtx           *gin.Context
	session          *session
	ormSession       *sorm.Session
	responseFinished bool
	firstPanicOffset int
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 最后一个方法为中间件时，说明没有匹配到url
		if strings.Contains(c.HandlerName(), "middleware") {
			return
		}

		session := newSession(c.Request, c.Writer)
		defer session.expiry()
		ormSession := sorm.NewSession(c.Request.Context())
		defer ormSession.Close()
		ctrl := &Controller{ginCtx: c, ormSession: ormSession, session: session}
		defer func() {
			if v := recover(); v != nil {
				if err, ok := v.(error); ok {
					for _, re := range []error{
						exception.RootError,
						exception.UnexpectedError,
					} {
						if errors.Is(err, re) {
							if werr, ok := err.(exception.ErrorWrapper); ok {
								if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut {
									ctrl.JsonResponseWithMsg(nil, werr.Error(), werr.Code())
								} else {
									ctrl.Echo(fmt.Sprintf("%d:%s", werr.Code(), werr.Error()))
								}
								logging.ErrorWithPrefix(werr.Position(), werr)
							}
							return
						}
					}
				}
				ctrl.Echo(fmt.Sprintf("%v", v))
			}
		}()
		c.Set(DefaultKey, ctrl)
		c.Next()
		ctrl.JsonResponse(nil)
	}
}

func Default(c *gin.Context) *Controller {
	return c.MustGet(DefaultKey).(*Controller)
}

func (ctrl *Controller) GetOrmSession() *sorm.Session {
	return ctrl.ormSession
}

func (ctrl *Controller) GetPar(key string) validator.ValueInterface {
	if v, ok := ctrl.ginCtx.Params.Get(key); ok {
		return validator.NewValue(v)
	}
	return validator.NewNil(true)
}

func (ctrl *Controller) GetQuery(key string, must bool) validator.ValueInterface {
	if v, ok := ctrl.ginCtx.GetQuery(key); ok {
		return validator.NewValue(v)
	}
	return validator.NewNil(must)
}

func (ctrl *Controller) DefaultQuery(key string, defaultValue string) validator.ValueInterface {
	if v, ok := ctrl.ginCtx.GetQuery(key); ok {
		return validator.NewValue(v)
	}
	return validator.NewValue(defaultValue).NoValidate()
}

func (ctrl *Controller) GetForm(key string, must bool) validator.ValueInterface {
	if v, ok := ctrl.ginCtx.GetPostForm(key); ok {
		return validator.NewValue(v)
	}
	return validator.NewNil(must)
}

func (ctrl *Controller) DefaultForm(key string, defaultValue string) validator.ValueInterface {
	if v, ok := ctrl.ginCtx.GetPostForm(key); ok {
		return validator.NewValue(v)
	}
	return validator.NewValue(defaultValue).NoValidate()
}

func (ctrl *Controller) ParsePost(v interface{}) {
	binding.Validator = validator.NewValidator()
	err := ctrl.ginCtx.Bind(v)
	if err != nil {
		exception.ThrowMsgWithCallerDepth(err.Error(), excode.ValidateError, 3)
	}
}

func (ctrl *Controller) JsonResponse(data interface{}) {
	if ctrl.responseFinished {
		return
	}
	ctrl.responseFinished = true
	response := util.JsonEncode(Response{
		Data:    data,
		Message: "ok",
		Code:    200,
	})
	EchoResponse(ctrl.ginCtx.Writer, response)
}

func (ctrl *Controller) JsonResponseWithMsg(data interface{}, message string, code int) {
	if ctrl.responseFinished {
		return
	}
	ctrl.EndRequest()
	response := util.JsonEncode(Response{
		Data:    data,
		Message: message,
		Code:    code,
	})
	EchoResponse(ctrl.ginCtx.Writer, response)
}

func (ctrl *Controller) EndRequest() {
	ctrl.responseFinished = true
}

func (ctrl *Controller) Echo(response string) {
	ctrl.EndRequest()
	EchoResponse(ctrl.ginCtx.Writer, response)
}

func (ctrl *Controller) GetSession() *session {
	return ctrl.session
}

func EchoResponse(writer http.ResponseWriter, response string) {
	_, _ = writer.Write([]byte(response))
}
