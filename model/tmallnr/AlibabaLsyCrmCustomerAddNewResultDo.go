package tmallnr

import (
	"sync"
)

// AlibabaLsyCrmCustomerAddNewResultDo 结构体
type AlibabaLsyCrmCustomerAddNewResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 出参对象
	Data *NrtRecordDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaLsyCrmCustomerAddNewResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaLsyCrmCustomerAddNewResultDo)
	},
}

// GetAlibabaLsyCrmCustomerAddNewResultDo() 从对象池中获取AlibabaLsyCrmCustomerAddNewResultDo
func GetAlibabaLsyCrmCustomerAddNewResultDo() *AlibabaLsyCrmCustomerAddNewResultDo {
	return poolAlibabaLsyCrmCustomerAddNewResultDo.Get().(*AlibabaLsyCrmCustomerAddNewResultDo)
}

// ReleaseAlibabaLsyCrmCustomerAddNewResultDo 释放AlibabaLsyCrmCustomerAddNewResultDo
func ReleaseAlibabaLsyCrmCustomerAddNewResultDo(v *AlibabaLsyCrmCustomerAddNewResultDo) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaLsyCrmCustomerAddNewResultDo.Put(v)
}
