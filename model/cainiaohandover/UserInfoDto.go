package cainiaohandover

import (
	"sync"
)

// UserInfoDto 结构体
type UserInfoDto struct {
	// 国家编码(选填)
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 登陆账号(必填)
	LoginId string `json:"login_id,omitempty" xml:"login_id,omitempty"`
	// 商家id(选填)
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 对应创建物流单的时候传入的top_user_key;跨店铺场景需要传入
	TopUserKey string `json:"top_user_key,omitempty" xml:"top_user_key,omitempty"`
	// 业务类型(选填)
	BizSource string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
	// 用户Key(选填)
	AppUserKey string `json:"app_user_key,omitempty" xml:"app_user_key,omitempty"`
	// 用户id(选填)
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolUserInfoDto = sync.Pool{
	New: func() any {
		return new(UserInfoDto)
	},
}

// GetUserInfoDto() 从对象池中获取UserInfoDto
func GetUserInfoDto() *UserInfoDto {
	return poolUserInfoDto.Get().(*UserInfoDto)
}

// ReleaseUserInfoDto 释放UserInfoDto
func ReleaseUserInfoDto(v *UserInfoDto) {
	v.Country = ""
	v.LoginId = ""
	v.SellerId = ""
	v.TopUserKey = ""
	v.BizSource = ""
	v.AppUserKey = ""
	v.UserId = ""
	poolUserInfoDto.Put(v)
}
