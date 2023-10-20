package topoaid

import (
	"sync"
)

// UserAuthInfo 结构体
type UserAuthInfo struct {
	// 收件人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 授权信息过期时间
	AuthorizeExpireTime string `json:"authorize_expire_time,omitempty" xml:"authorize_expire_time,omitempty"`
	// 用户唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 虚拟号
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 虚拟号过期时间
	SecretExpireTime string `json:"secret_expire_time,omitempty" xml:"secret_expire_time,omitempty"`
	// 运单所属快递公司code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 收件人手机号是否为虚拟号主号(虚拟号前11位)
	MainSecretNo bool `json:"main_secret_no,omitempty" xml:"main_secret_no,omitempty"`
	// 收件人手机号是否在柜机黑名单中
	Black bool `json:"black,omitempty" xml:"black,omitempty"`
}

var poolUserAuthInfo = sync.Pool{
	New: func() any {
		return new(UserAuthInfo)
	},
}

// GetUserAuthInfo() 从对象池中获取UserAuthInfo
func GetUserAuthInfo() *UserAuthInfo {
	return poolUserAuthInfo.Get().(*UserAuthInfo)
}

// ReleaseUserAuthInfo 释放UserAuthInfo
func ReleaseUserAuthInfo(v *UserAuthInfo) {
	v.Mobile = ""
	v.AuthorizeExpireTime = ""
	v.OpenId = ""
	v.SecretNo = ""
	v.SecretExpireTime = ""
	v.CpCode = ""
	v.MainSecretNo = false
	v.Black = false
	poolUserAuthInfo.Put(v)
}
