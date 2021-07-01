package miniappopen

// MiniAppEntityDto 结构体
type MiniAppEntityDto struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 小程序appId
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 小程序名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 小程序描述
	AppDescription string `json:"app_description,omitempty" xml:"app_description,omitempty"`
	// icon
	AppIcon string `json:"app_icon,omitempty" xml:"app_icon,omitempty"`
	// 当前版本
	OnlineVersion string `json:"online_version,omitempty" xml:"online_version,omitempty"`
	// 线上码
	OnlineCode string `json:"online_code,omitempty" xml:"online_code,omitempty"`
}
