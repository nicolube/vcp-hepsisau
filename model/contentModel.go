package model

type ContentModel struct {
	Model
	UserId    int64
	Type      string
	Content   string
	CreatedAt UnixTimeS
}
