//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package adminForm

import (
	"github.com/bufanyun/hotgo/app/form"
	"github.com/bufanyun/hotgo/app/form/input"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/bufanyun/hotgo/app/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

//  更新会员资料
type MemberUpdateProfileReq struct {
	g.Meta   `path:"/member/update_profile" method:"post" tags:"会员" summary:"更新会员资料"`
	Mobile   int    `json:"mobile"   dc:"手机号"`
	Email    string `json:"email"   dc:"邮箱"`
	Realname string `json:"realname"   dc:"真实姓名"`
}
type MemberUpdateProfileRes struct{}

//  修改登录密码
type MemberUpdatePwdReq struct {
	g.Meta      `path:"/member/update_pwd" method:"post" tags:"会员" summary:"重置密码"`
	OldPassword string `json:"oldPassword" v:"required#原密码不能为空"  dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,16#新密码不能为空#新密码需在6~16之间"  dc:"新密码"`
}
type MemberUpdatePwdRes struct{}

//  获取登录用户的基本信息
type MemberProfileReq struct {
	g.Meta `path:"/member/profile" method:"get" tags:"会员" summary:"获取登录用户的基本信息"`
}
type MemberProfileRes struct {
	PostGroup string                      `json:"postGroup" dc:"岗位名称"`
	RoleGroup string                      `json:"roleGroup" dc:"角色名称"`
	User      *input.AdminMemberViewModel `json:"user" dc:"用户基本信息"`
	SysDept   *input.AdminDeptViewModel   `json:"sysDept" dc:"部门信息"`
	SysRoles  []*input.AdminRoleListModel `json:"sysRoles" dc:"角色列表"`
	PostIds   int64                       `json:"postIds" dc:"当前岗位"`
	RoleIds   int64                       `json:"roleIds" dc:"当前角色"`
}

//  重置密码
type MemberResetPwdReq struct {
	g.Meta   `path:"/member/reset_pwd" method:"post" tags:"会员" summary:"重置密码"`
	Password string `json:"password" v:"required#密码不能为空"  dc:"密码"`
	Id       int64  `json:"id" dc:"会员ID"`
}
type MemberResetPwdRes struct{}

//  邮箱是否唯一
type MemberEmailUniqueReq struct {
	g.Meta `path:"/member/email_unique" method:"get" tags:"会员" summary:"邮箱是否唯一"`
	Email  string `json:"email" v:"required#邮箱不能为空"  dc:"邮箱"`
	Id     int64  `json:"id" dc:"会员ID"`
}
type MemberEmailUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  手机号是否唯一
type MemberMobileUniqueReq struct {
	g.Meta `path:"/member/mobile_unique" method:"get" tags:"会员" summary:"手机号是否唯一"`
	Mobile string `json:"mobile" v:"required#手机号不能为空"  dc:"手机号"`
	Id     int64  `json:"id" dc:"会员ID"`
}
type MemberMobileUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  名称是否唯一
type MemberNameUniqueReq struct {
	g.Meta   `path:"/member/name_unique" method:"get" tags:"会员" summary:"会员名称是否唯一"`
	Username string `json:"username" v:"required#会员名称不能为空"  dc:"会员名称"`
	Id       int64  `json:"id" dc:"会员ID"`
}
type MemberNameUniqueRes struct {
	IsUnique bool `json:"is_unique" dc:"是否唯一"`
}

//  查询列表
type MemberListReq struct {
	g.Meta `path:"/member/list" method:"get" tags:"会员" summary:"获取会员列表"`
	form.PageReq
	form.RangeDateReq
	form.StatusReq
	DeptId    int    `json:"dept_id"   dc:"部门ID"`
	Mobile    int    `json:"mobile"   dc:"手机号"`
	Username  string `json:"username"   dc:"用户名"`
	Realname  string `json:"realname"   dc:"真实姓名"`
	StartTime string `json:"start_time"   dc:"开始时间"`
	EndTime   string `json:"end_time"   dc:"结束时间"`
	Name      string `json:"name"   dc:"岗位名称"`
	Code      string `json:"code"   dc:"岗位编码"`
}

type MemberListRes struct {
	List []*input.AdminMemberListModel `json:"list"   dc:"数据列表"`
	form.PageRes
}

//  获取指定信息
type MemberViewReq struct {
	g.Meta `path:"/member/view" method:"get" tags:"会员" summary:"获取指定信息"`
	Id     int64 `json:"id" dc:"会员ID"` // v:"required#会员ID不能为空"
}
type MemberViewRes struct {
	*input.AdminMemberViewModel
	Posts    []*input.AdminPostListModel `json:"posts" dc:"可选岗位"`
	PostIds  []int64                     `json:"postIds" dc:"当前岗位"`
	Roles    []*input.AdminRoleListModel `json:"roles" dc:"可选角色"`
	RoleIds  []int64                     `json:"roleIds" dc:"当前角色"`
	DeptName string                      `json:"dept_name" dc:"部门名称"`
}

//  修改/新增
type MemberEditReq struct {
	g.Meta `path:"/member/edit" method:"post" tags:"会员" summary:"修改/新增会员"`
	input.AdminMemberEditInp
}
type MemberEditRes struct{}

//  删除
type MemberDeleteReq struct {
	g.Meta `path:"/member/delete" method:"post" tags:"会员" summary:"删除会员"`
	Id     interface{} `json:"id" v:"required#会员ID不能为空" dc:"会员ID"`
}
type MemberDeleteRes struct{}

//  最大排序
type MemberMaxSortReq struct {
	g.Meta `path:"/member/max_sort" method:"get" tags:"会员" summary:"会员最大排序"`
	Id     int64 `json:"id" dc:"会员ID"`
}
type MemberMaxSortRes struct {
	Sort int `json:"sort" dc:"排序"`
}

//  获取登录用户信息
type MemberInfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"会员" summary:"获取登录用户信息" dc:"获取管理后台的登录用户信息"`
}

type PortalConfigContentOptions struct {
	TitleRequired bool   `json:"titleRequired"  titleRequired:""`
	MoreUrl       string `json:"moreUrl"  dc:"模块地址"`
	Refresh       int    `json:"refresh"  dc:"刷新"`
}

type PortalConfigContent struct {
	Id          int                           `json:"id"  dc:"内容ID"`
	X           int                           `json:"x"  dc:""`
	Y           int                           `json:"y"  dc:""`
	W           int                           `json:"w"  dc:"宽"`
	H           int                           `json:"h"  dc:"高"`
	I           int                           `json:"i"  dc:""`
	Key         string                        `json:"key"  dc:""`
	IsShowTitle string                        `json:"isShowTitle"  dc:""`
	IsAllowDrag bool                          `json:"isAllowDrag"  dc:""`
	Name        string                        `json:"name"  dc:""`
	Type        string                        `json:"type"  dc:""`
	Url         string                        `json:"url"  dc:""`
	Options     []*PortalConfigContentOptions `json:"options"  dc:""`
	Moved       bool                          `json:"moved"  dc:""`
}

type PortalConfig struct {
	CreateByName        string      `json:"createByName"  dc:"创建者名称"`
	CreateDeptName      string      `json:"createDeptName"  dc:"创建部门名称"`
	ImportErrInfo       string      `json:"importErrInfo"  dc:"导出错误信息"`
	Id                  string      `json:"id"  dc:"用户ID"`
	SearchValue         string      `json:"searchValue"  dc:"搜索内容"`
	CreateBy            string      `json:"createBy"  dc:"创建者名称"`
	CreateDept          string      `json:"createDept"  dc:"创建部门名称"`
	CreateTime          *gtime.Time `json:"createTime"  dc:"创建时间"`
	UpdateBy            string      `json:"updateBy"  dc:"更新者名称"`
	UpdateTime          *gtime.Time `json:"updateTime"  dc:"更新时间"`
	UpdateIp            string      `json:"updateIp"  dc:"更新iP"`
	Remark              string      `json:"remark"  dc:"备注"`
	Version             string      `json:"version"  dc:"版本号"`
	DelFlag             string      `json:"delFlag"  dc:"删除标签"`
	HandleType          string      `json:"handleType"  dc:""`
	Params              string      `json:"params"  dc:""`
	Name                string      `json:"name"  dc:"配置名称"`
	Code                string      `json:"code"  dc:"配置代码"`
	ApplicationRange    string      `json:"applicationRange"  dc:""`
	IsDefault           string      `json:"isDefault"  dc:"是否默认"`
	ResourceId          string      `json:"resourceId"  dc:""`
	ResourceName        string      `json:"resourceName"  dc:""`
	SystemDefinedId     string      `json:"systemDefinedId"  dc:""`
	Sort                string      `json:"sort"  dc:"排序"`
	SaveType            string      `json:"saveType"  dc:""`
	Status              string      `json:"status"  dc:"状态"`
	RecordLog           string      `json:"recordLog"  dc:""`
	PortalConfigContent string      `json:"content"  dc:"配置内容"`
}

type MemberInfoRes struct {
	DefaultPortalConfig []*PortalConfig       `json:"defaultPortalConfig"  dc:"默认用户配置"`
	LincenseInfo        string                `json:"lincenseInfo"  dc:"应用版本号"`
	Permissions         []string              `json:"permissions"dc:"权限"`
	Roles               []string              `json:"roles" dc:"角色"`
	SysNoticeList       []*entity.AdminNotice `json:"sysNoticeList" dc:"系统公告"`
	UserPortalConfig    []*PortalConfig       `json:"userPortalConfig" dc:"用户配置"`
	User                model.Identity        `json:"user"  dc:"用户信息"`
}
