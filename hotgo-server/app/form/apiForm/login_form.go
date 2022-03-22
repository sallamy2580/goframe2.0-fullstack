//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiForm

import (
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/frame/g"
)

//  注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/login/logout" method:"get" tags:"登录接口" summary:"注销登录"`
}
type LoginLogoutRes struct{}

//  登录效验
type LoginCheckReq struct {
	g.Meta `path:"/login/check" method:"get" tags:"登录接口" summary:"登录效验"`
}
type LoginCheckRes struct {
	IsValidCodeLogin bool   `json:"isValidCodeLogin"      dc:"是否验证码"`
	Message          string `json:"message"      dc:"消息"`
	Result           string `json:"result"      dc:"响应"`
	// Sessionid        string `json:"sessionid"      dc:"sessionid"`
}

//  提交登录
type LoginReq struct {
	g.Meta   `path:"/login/sign" method:"post" tags:"登录接口" summary:"提交登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Cid      string `json:"cid" v:"required#验证码ID不能为空" dc:"验证码ID"`
	Code     string `json:"code" v:"required#验证码不能为空" dc:"验证码"`
	Device   string `json:"device"  dc:"登录设备"`
}
type LoginRes struct {
	model.Identity
	Token string `json:"token" dc:"登录token"`
}
