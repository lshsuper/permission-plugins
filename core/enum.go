package core

type MenuType int

const(
	All MenuType=iota
	Dir
	Link
	Btn
)


//Remark
func (e MenuType)Remark()string  {
	switch e {
	case Dir:
		return "dir"
	case Link:
		return "link"
	case Btn:
		return "btn"
	default:
		return ""
	}
}

//GetDescription
func (e MenuType)GetDescription()string  {
	switch e {
	case Dir:
		return "目录"
	case Link:
		return "链接"
	case Btn:
		return "按钮"
	default:
		return ""
	}
}
