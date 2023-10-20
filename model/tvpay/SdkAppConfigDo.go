package tvpay

import (
	"sync"
)

// SdkAppConfigDo 结构体
type SdkAppConfigDo struct {
	// 自定义属性
	ConfigProps string `json:"config_props,omitempty" xml:"config_props,omitempty"`
	// 是否上传日志
	EnableUploadLog bool `json:"enable_upload_log,omitempty" xml:"enable_upload_log,omitempty"`
	// 是否开启友盟
	EnableYoument bool `json:"enable_youment,omitempty" xml:"enable_youment,omitempty"`
	// 是否自动登入
	EnableAutoLogin bool `json:"enable_auto_login,omitempty" xml:"enable_auto_login,omitempty"`
}

var poolSdkAppConfigDo = sync.Pool{
	New: func() any {
		return new(SdkAppConfigDo)
	},
}

// GetSdkAppConfigDo() 从对象池中获取SdkAppConfigDo
func GetSdkAppConfigDo() *SdkAppConfigDo {
	return poolSdkAppConfigDo.Get().(*SdkAppConfigDo)
}

// ReleaseSdkAppConfigDo 释放SdkAppConfigDo
func ReleaseSdkAppConfigDo(v *SdkAppConfigDo) {
	v.ConfigProps = ""
	v.EnableUploadLog = false
	v.EnableYoument = false
	v.EnableAutoLogin = false
	poolSdkAppConfigDo.Put(v)
}
