package controllers

import (
	"decoration-admin/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// UserController handles user related requests
type UserController struct {
	beego.Controller
}

type ResBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (u *UserController) unmarshalPayload(v interface{}) error {
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, v); err != nil {
		logs.Error("unmarshal payload of %s error: %s", u.Ctx.Request.URL.Path, err)
	}
	return nil
}

func (u *UserController) respond(code int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(code)
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	u.Data["json"] = ResBody{
		Code:    code,
		Message: message,
		Data:    d,
	}
	u.ServeJSON()
}

// Login handles login request
func (u *UserController) Login() {
	lr := new(models.LoginRequest)
	if err := u.unmarshalPayload(lr); err != nil {
		u.respond(http.StatusBadRequest, err.Error())
		return
	}
	lrs, statusCode, err := models.DoLogin(lr)
	if err != nil {
		u.respond(statusCode, err.Error())
		return
	}
	//u.Ctx.Output.Header("Authorization", lrs.Token) // set token into header
	u.respond(http.StatusOK, "", lrs)
}

// Get user`s info
func (u *UserController) GetUserInfo() {
	logs.Info("Get user`s info")

}

// CreateUser creates a user
func (u *UserController) CreateUser() {
	cu := new(models.CreateRequest)
	if err := u.unmarshalPayload(cu); err != nil {
		u.respond(http.StatusBadRequest, err.Error())
	}
	createUser, statusCode, err := models.DoCreateUser(cu)
	if err != nil {
		u.respond(statusCode, err.Error())
		return
	}
	u.respond(http.StatusOK, "", createUser)
}

// Valid token
func TokenAuth(ctx *context.Context) {
	token := ctx.Input.Header("X-Token")
	logs.Info("Valid token", token)
	if token != "" {
		var err error
		_, err = models.ValidateToken(token)
		if err != nil {
			ctx.Output.JSON(ResBody{Code: http.StatusBadRequest, Message: "验证失败", Data: ""}, true, true)
		} else {
			return
		}
	} else {
		ctx.Output.JSON(ResBody{Code: http.StatusBadRequest, Message: "验证失败", Data: ""}, true, true)
	}
}
