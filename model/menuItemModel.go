package model

type MenuItemModel struct {
	Model
	ParentId *int64
	Children []MenuItemModel
	Name     string
	SortId   int64
	Side     *SideModel
}
