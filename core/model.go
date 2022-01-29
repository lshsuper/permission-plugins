package core

type Roles struct {
	ID int `json:"id" gorm:"primary_key column:id"`
	RoleName string `json:"role_name" gorm:"role_name"`
}


type Menus struct {
	ID int `json:"id" gorm:"primary_key column:id"`
	MenuName string `json:"menu_name" gorm:"column:menu_name"`
	MenuType MenuType `json:"menu_type" gorm:"column:menu_type"`
}
