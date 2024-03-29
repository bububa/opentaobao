package alitripmerchant

import (
	"sync"
)

// LoginParam 结构体
type LoginParam struct {
	// 用户信息初始化向量
	InfoIv string `json:"info_iv,omitempty" xml:"info_iv,omitempty"`
	// 微信认证code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 加密的手机信息
	EncryptedPhone string `json:"encrypted_phone,omitempty" xml:"encrypted_phone,omitempty"`
	// 验签数据完整签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 用户ip
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 用户信息的加密数据
	EncryptedInfo string `json:"encrypted_info,omitempty" xml:"encrypted_info,omitempty"`
	// 解密用户手机号的初始化向量
	PhoneIv string `json:"phone_iv,omitempty" xml:"phone_iv,omitempty"`
	// 用户的基本信息
	RawData string `json:"raw_data,omitempty" xml:"raw_data,omitempty"`
	// 场景值
	SceneDistinction string `json:"scene_distinction,omitempty" xml:"scene_distinction,omitempty"`
	// 姓名验证
	NameVerification string `json:"name_verification,omitempty" xml:"name_verification,omitempty"`
	// 用户信息加密版
	UserEncryptInfo string `json:"user_encrypt_info,omitempty" xml:"user_encrypt_info,omitempty"`
	// 用户注册来源
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 新版本用户信息
	NewUserinfo *NewUserInfo `json:"new_userinfo,omitempty" xml:"new_userinfo,omitempty"`
	// 是否是老版本
	OldVersion bool `json:"old_version,omitempty" xml:"old_version,omitempty"`
	// 注销账号场景
	DestroyVerification bool `json:"destroy_verification,omitempty" xml:"destroy_verification,omitempty"`
}

var poolLoginParam = sync.Pool{
	New: func() any {
		return new(LoginParam)
	},
}

// GetLoginParam() 从对象池中获取LoginParam
func GetLoginParam() *LoginParam {
	return poolLoginParam.Get().(*LoginParam)
}

// ReleaseLoginParam 释放LoginParam
func ReleaseLoginParam(v *LoginParam) {
	v.InfoIv = ""
	v.Code = ""
	v.EncryptedPhone = ""
	v.Signature = ""
	v.Ip = ""
	v.EncryptedInfo = ""
	v.PhoneIv = ""
	v.RawData = ""
	v.SceneDistinction = ""
	v.NameVerification = ""
	v.UserEncryptInfo = ""
	v.Channel = 0
	v.NewUserinfo = nil
	v.OldVersion = false
	v.DestroyVerification = false
	poolLoginParam.Put(v)
}
