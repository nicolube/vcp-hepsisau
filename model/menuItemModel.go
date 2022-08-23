package model

type MenuItemModel struct {
	Model
	Children []MenuItemModel
	Name     string
	SortId   int64
	Side     SideModel
}
