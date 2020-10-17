package db

type User struct {
	Id         int32  `db:"id" json:"id"`
	SessionKey string `db:"session_key" json:"session_key"`
	Openid     string `db:"openid" json:"openid"`
}
