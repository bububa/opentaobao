package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmCustomerAddResultDo 结构体
type AlibabaLsyCrmCustomerAddResultDo struct {
	// 错误提示码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmCustomerAddResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmCustomerAddResultDo)
	},
}

// GetAlibabaLsyCrmCustomerAddResultDo() 从对象池中获取AlibabaLsyCrmCustomerAddResultDo
func GetAlibabaLsyCrmCustomerAddResultDo() *AlibabaLsyCrmCustomerAddResultDo {
	return poolAlibabaLsyCrmCustomerAddResultDo.Get().(*AlibabaLsyCrmCustomerAddResultDo)
}

// ReleaseAlibabaLsyCrmCustomerAddResultDo 释放AlibabaLsyCrmCustomerAddResultDo
func ReleaseAlibabaLsyCrmCustomerAddResultDo(v *AlibabaLsyCrmCustomerAddResultDo) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaLsyCrmCustomerAddResultDo.Put(v)
}
