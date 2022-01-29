package core

import "gorm.io/gorm"

type permissionProvider struct {

}


type Conf struct {

}
//Register 注册插件
func Register(conf Conf)*permissionProvider  {
	return &permissionProvider{}
}
//GetRoles 所有角色
func (provider *permissionProvider)GetRoles(db *gorm.DB)(lst []Roles)  {
	db.Find(&lst,"del=0")
    return
}

//GetMenus 所有菜单
func (provider *permissionProvider)GetMenus(db *gorm.DB)(lst []Menus)  {
	db.Find(&lst,"del=0")
	return
}





