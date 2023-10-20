package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmActivityDataUpdateResultDo 结构体
type AlibabaLsyCrmActivityDataUpdateResultDo struct {
	// err_code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmActivityDataUpdateResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityDataUpdateResultDo)
	},
}

// GetAlibabaLsyCrmActivityDataUpdateResultDo() 从对象池中获取AlibabaLsyCrmActivityDataUpdateResultDo
func GetAlibabaLsyCrmActivityDataUpdateResultDo() *AlibabaLsyCrmActivityDataUpdateResultDo {
	return poolAlibabaLsyCrmActivityDataUpdateResultDo.Get().(*AlibabaLsyCrmActivityDataUpdateResultDo)
}

// ReleaseAlibabaLsyCrmActivityDataUpdateResultDo 释放AlibabaLsyCrmActivityDataUpdateResultDo
func ReleaseAlibabaLsyCrmActivityDataUpdateResultDo(v *AlibabaLsyCrmActivityDataUpdateResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaLsyCrmActivityDataUpdateResultDo.Put(v)
}
