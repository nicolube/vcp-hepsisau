package model

type UserTokenModel struct {
	Model
	UserId     int
	Token      string
	Ip         string
	LastUsedAt UnixTimeS
	CreatedAt  UnixTimeS
}
