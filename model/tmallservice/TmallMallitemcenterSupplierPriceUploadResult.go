package tmallservice

import (
	"sync"
)

// TmallMallitemcenterSupplierPriceUploadResult 结构体
type TmallMallitemcenterSupplierPriceUploadResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	ResultData *ResultData `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否系统调用错误
	SystemError bool `json:"system_error,omitempty" xml:"system_error,omitempty"`
	// 是否校验出错
	BusinessCheckFail bool `json:"business_check_fail,omitempty" xml:"business_check_fail,omitempty"`
	// 接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallMallitemcenterSupplierPriceUploadResult = sync.Pool{
	New: func() any {
		return new(TmallMallitemcenterSupplierPriceUploadResult)
	},
}

// GetTmallMallitemcenterSupplierPriceUploadResult() 从对象池中获取TmallMallitemcenterSupplierPriceUploadResult
func GetTmallMallitemcenterSupplierPriceUploadResult() *TmallMallitemcenterSupplierPriceUploadResult {
	return poolTmallMallitemcenterSupplierPriceUploadResult.Get().(*TmallMallitemcenterSupplierPriceUploadResult)
}

// ReleaseTmallMallitemcenterSupplierPriceUploadResult 释放TmallMallitemcenterSupplierPriceUploadResult
func ReleaseTmallMallitemcenterSupplierPriceUploadResult(v *TmallMallitemcenterSupplierPriceUploadResult) {
	v.Message = ""
	v.ErrorCode = ""
	v.ResultData = nil
	v.SystemError = false
	v.BusinessCheckFail = false
	v.Success = false
	poolTmallMallitemcenterSupplierPriceUploadResult.Put(v)
}
