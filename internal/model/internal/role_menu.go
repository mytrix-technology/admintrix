package internal

// RoleMenu is the golang structure for table sys_role_menu.
type RoleMenu struct {
	RoleId int `orm:"role_id" json:"roleId"` // 角色ID
	MenuId int `orm:"menu_id" json:"menuId"` // 菜单ID
}