package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmActivityGetbaseinfoResultDo 结构体
type AlibabaLsyCrmActivityGetbaseinfoResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// DTO
	Data *NrtCrmActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmActivityGetbaseinfoResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmActivityGetbaseinfoResultDo)
	},
}

// GetAlibabaLsyCrmActivityGetbaseinfoResultDo() 从对象池中获取AlibabaLsyCrmActivityGetbaseinfoResultDo
func GetAlibabaLsyCrmActivityGetbaseinfoResultDo() *AlibabaLsyCrmActivityGetbaseinfoResultDo {
	return poolAlibabaLsyCrmActivityGetbaseinfoResultDo.Get().(*AlibabaLsyCrmActivityGetbaseinfoResultDo)
}

// ReleaseAlibabaLsyCrmActivityGetbaseinfoResultDo 释放AlibabaLsyCrmActivityGetbaseinfoResultDo
func ReleaseAlibabaLsyCrmActivityGetbaseinfoResultDo(v *AlibabaLsyCrmActivityGetbaseinfoResultDo) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLsyCrmActivityGetbaseinfoResultDo.Put(v)
}
