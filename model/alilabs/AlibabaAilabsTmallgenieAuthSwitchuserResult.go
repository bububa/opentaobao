package alilabs

import (
	"sync"
)

// AlibabaAilabsTmallgenieAuthSwitchuserResult 结构体
type AlibabaAilabsTmallgenieAuthSwitchuserResult struct {
	// 扩展信息
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 注册结果
	RegisterInfoVO *RegisterInfoVo `json:"register_info_v_o,omitempty" xml:"register_info_v_o,omitempty"`
	// 授权结果
	DeviceTokenVO *DeviceTokenVo `json:"device_token_v_o,omitempty" xml:"device_token_v_o,omitempty"`
}

var poolAlibabaAilabsTmallgenieAuthSwitchuserResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthSwitchuserResult)
	},
}

// GetAlibabaAilabsTmallgenieAuthSwitchuserResult() 从对象池中获取AlibabaAilabsTmallgenieAuthSwitchuserResult
func GetAlibabaAilabsTmallgenieAuthSwitchuserResult() *AlibabaAilabsTmallgenieAuthSwitchuserResult {
	return poolAlibabaAilabsTmallgenieAuthSwitchuserResult.Get().(*AlibabaAilabsTmallgenieAuthSwitchuserResult)
}

// ReleaseAlibabaAilabsTmallgenieAuthSwitchuserResult 释放AlibabaAilabsTmallgenieAuthSwitchuserResult
func ReleaseAlibabaAilabsTmallgenieAuthSwitchuserResult(v *AlibabaAilabsTmallgenieAuthSwitchuserResult) {
	v.Extension = ""
	v.RegisterInfoVO = nil
	v.DeviceTokenVO = nil
	poolAlibabaAilabsTmallgenieAuthSwitchuserResult.Put(v)
}
