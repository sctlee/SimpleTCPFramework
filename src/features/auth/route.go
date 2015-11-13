package auth

import (
	. "features/auth/action"

	"github.com/sctlee/tcpx"
)

var userAction = NewUserAction()

var Router = map[string]tcpx.RouteFun{
	"setName": userAction.SetUserName,
	"login":   userAction.Login,
	"logout":  userAction.Logout,
	"signup":  userAction.Signup,
}