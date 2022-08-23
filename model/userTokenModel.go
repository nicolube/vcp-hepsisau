package model

type UserTokenModel struct {
	Model
	UserId     int64
	Token      string
	Ip         string
	LastUsedAt UnixTimeS
	CreatedAt  UnixTimeS
}
