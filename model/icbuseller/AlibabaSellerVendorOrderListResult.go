package icbuseller

import (
	"sync"
)

// AlibabaSellerVendorOrderListResult 结构体
type AlibabaSellerVendorOrderListResult struct {
	// 返回对象集合
	Dtos []Dto `json:"dtos,omitempty" xml:"dtos>dto,omitempty"`
	// 接口返回
	ExecDescription string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
	// 返回码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 分页对象
	PageDto *PageDto `json:"page_dto,omitempty" xml:"page_dto,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSellerVendorOrderListResult = sync.Pool{
	New: func() any {
		return new(AlibabaSellerVendorOrderListResult)
	},
}

// GetAlibabaSellerVendorOrderListResult() 从对象池中获取AlibabaSellerVendorOrderListResult
func GetAlibabaSellerVendorOrderListResult() *AlibabaSellerVendorOrderListResult {
	return poolAlibabaSellerVendorOrderListResult.Get().(*AlibabaSellerVendorOrderListResult)
}

// ReleaseAlibabaSellerVendorOrderListResult 释放AlibabaSellerVendorOrderListResult
func ReleaseAlibabaSellerVendorOrderListResult(v *AlibabaSellerVendorOrderListResult) {
	v.Dtos = v.Dtos[:0]
	v.ExecDescription = ""
	v.ReturnCode = 0
	v.PageDto = nil
	v.Success = false
	poolAlibabaSellerVendorOrderListResult.Put(v)
}
