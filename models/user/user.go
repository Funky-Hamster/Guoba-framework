package usermodel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"guoba.io/guobastreet/util"
)

// User is
type User struct {
	Id         int32  `db:"id" json:"id"`
	SessionKey string `db:"session_key" json:"session_key"`
	Openid     string `db:"openid" json:"openid"`
}

type UserOperation struct {
}

func NewUserOperation() *UserOperation {
	return &UserOperation{}
}

type code2SessionResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Code       int    `json:"errcode"`
	Msg        string `json:"errmsg"`
}

func (u *UserOperation) GetUserByOpenid(jsCode string) (*User, error) {
	var user *User = &User{}
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + util.WX_APPID + "&secret=" + util.WX_SECRET + "&js_code=" + jsCode + "&grant_type=authorization_code")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res code2SessionResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	user.Openid = res.Openid
	user.SessionKey = res.SessionKey

	if res.Code != 0 {
		return nil, nil
	}

	return user, nil

}
