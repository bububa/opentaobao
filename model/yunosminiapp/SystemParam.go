package yunosminiapp

import (
	"sync"
)

// SystemParam 结构体
type SystemParam struct {
	// 流程id，随机字符串
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 业务码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 小程序id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 授权token
	AcToken string `json:"ac_token,omitempty" xml:"ac_token,omitempty"`
	// 设备信息
	DeviceInfo string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 更新access_token
	ModifyToken bool `json:"modify_token,omitempty" xml:"modify_token,omitempty"`
	// token过期
	TokenExpired bool `json:"token_expired,omitempty" xml:"token_expired,omitempty"`
}

var poolSystemParam = sync.Pool{
	New: func() any {
		return new(SystemParam)
	},
}

// GetSystemParam() 从对象池中获取SystemParam
func GetSystemParam() *SystemParam {
	return poolSystemParam.Get().(*SystemParam)
}

// ReleaseSystemParam 释放SystemParam
func ReleaseSystemParam(v *SystemParam) {
	v.TraceId = ""
	v.BizCode = ""
	v.AppId = ""
	v.AcToken = ""
	v.DeviceInfo = ""
	v.ModifyToken = false
	v.TokenExpired = false
	poolSystemParam.Put(v)
}
