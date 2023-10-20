package tvpay

import (
	"sync"
)

// SdkGlobalConfigDo 结构体
type SdkGlobalConfigDo struct {
	// h5路由属性
	H5ProxyUrl string `json:"h5_proxy_url,omitempty" xml:"h5_proxy_url,omitempty"`
	// 询问对话框样式
	AutoLoginDialogPattern int64 `json:"auto_login_dialog_pattern,omitempty" xml:"auto_login_dialog_pattern,omitempty"`
	// 是否询问用户要自动登录
	AskBeforeAutoLogin bool `json:"ask_before_auto_login,omitempty" xml:"ask_before_auto_login,omitempty"`
	// 如果询问，是否默认选中自动登录
	AutoLoginChecked bool `json:"auto_login_checked,omitempty" xml:"auto_login_checked,omitempty"`
	// 是否默认展示授权二维码
	ShowAuthCodeByDefault bool `json:"show_auth_code_by_default,omitempty" xml:"show_auth_code_by_default,omitempty"`
}

var poolSdkGlobalConfigDo = sync.Pool{
	New: func() any {
		return new(SdkGlobalConfigDo)
	},
}

// GetSdkGlobalConfigDo() 从对象池中获取SdkGlobalConfigDo
func GetSdkGlobalConfigDo() *SdkGlobalConfigDo {
	return poolSdkGlobalConfigDo.Get().(*SdkGlobalConfigDo)
}

// ReleaseSdkGlobalConfigDo 释放SdkGlobalConfigDo
func ReleaseSdkGlobalConfigDo(v *SdkGlobalConfigDo) {
	v.H5ProxyUrl = ""
	v.AutoLoginDialogPattern = 0
	v.AskBeforeAutoLogin = false
	v.AutoLoginChecked = false
	v.ShowAuthCodeByDefault = false
	poolSdkGlobalConfigDo.Put(v)
}
