package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmActivityGetResultDo 结构体
type AlibabaLsyCrmActivityGetResultDo struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 数据
	Data *NrtCrmActivityDetailDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmActivityGetResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityGetResultDo)
	},
}

// GetAlibabaLsyCrmActivityGetResultDo() 从对象池中获取AlibabaLsyCrmActivityGetResultDo
func GetAlibabaLsyCrmActivityGetResultDo() *AlibabaLsyCrmActivityGetResultDo {
	return poolAlibabaLsyCrmActivityGetResultDo.Get().(*AlibabaLsyCrmActivityGetResultDo)
}

// ReleaseAlibabaLsyCrmActivityGetResultDo 释放AlibabaLsyCrmActivityGetResultDo
func ReleaseAlibabaLsyCrmActivityGetResultDo(v *AlibabaLsyCrmActivityGetResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLsyCrmActivityGetResultDo.Put(v)
}
