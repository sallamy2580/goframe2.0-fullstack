// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
)

// adminMenuOldDao is the data access object for table hg_admin_menu_old.
// You can define custom methods on it to extend its functionality as you wish.
type adminMenuOldDao struct {
	*internal.AdminMenuOldDao
}

var (
	// AdminMenuOld is globally public accessible object for table hg_admin_menu_old operations.
	AdminMenuOld = adminMenuOldDao{
		internal.NewAdminMenuOldDao(),
	}
)

// Fill with you ideas below.
