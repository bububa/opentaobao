package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmActivityUpdateResultDo 结构体
type AlibabaLsyCrmActivityUpdateResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmActivityUpdateResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityUpdateResultDo)
	},
}

// GetAlibabaLsyCrmActivityUpdateResultDo() 从对象池中获取AlibabaLsyCrmActivityUpdateResultDo
func GetAlibabaLsyCrmActivityUpdateResultDo() *AlibabaLsyCrmActivityUpdateResultDo {
	return poolAlibabaLsyCrmActivityUpdateResultDo.Get().(*AlibabaLsyCrmActivityUpdateResultDo)
}

// ReleaseAlibabaLsyCrmActivityUpdateResultDo 释放AlibabaLsyCrmActivityUpdateResultDo
func ReleaseAlibabaLsyCrmActivityUpdateResultDo(v *AlibabaLsyCrmActivityUpdateResultDo) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaLsyCrmActivityUpdateResultDo.Put(v)
}
