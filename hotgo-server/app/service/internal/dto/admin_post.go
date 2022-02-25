// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPost is the golang structure of table hg_admin_post for DAO operations like Where/Data.
type AdminPost struct {
	g.Meta    `orm:"table:hg_admin_post, dto:true"`
	Id        interface{} // 岗位ID
	Code      interface{} // 岗位编码
	Name      interface{} // 岗位名称
	Remark    interface{} // 备注
	Sort      interface{} // 显示顺序
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
