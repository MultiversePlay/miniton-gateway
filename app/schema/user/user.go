package user

import (
	"context"
	"errors"
)

const (
	HeaderAuthKey = "_auth"
)

const (
	GenderNotSet = -1
	GenderMan    = 1
	GenderWomen  = 2
)

const (
	// 默认头像
	DefaultAvatar = "default_avatar.jpg"
	// 默认背景图
	DefaultBackgroud = "default_backgroud.jpg"
)

type (
	TGUserKey struct{}
)

type (
	DetailReq struct {
		TGUserInfo *TGUserInfo
	}

	DetailResp struct {
		TGID       int64  `json:"tg_id"`
		UserID     int64  `json:"user_id"`
		Name       string `json:"name"`
		Region     int64  `json:"region"`
		Gender     int    `json:"gender"`
		Birth      int    `json:"birth"`
		Avatar     string `json:"avatar"`
		Backgroud  string `json:"backgroud"`
		TonAddress string `json:"ton_address"`
		TonBalance string `json:"ton_balance"`
	}
)

type (
	TGUserInfo struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
		PhotoUrl     string `json:"photo_url"`
	}
)

func GetUserInfo(c context.Context) (*TGUserInfo, error) {
	tgUserInfo := c.Value(TGUserKey{})
	t, ok := tgUserInfo.(*TGUserInfo)
	if !ok {
		return nil, errors.New("user err")
	}
	return t, nil
}
