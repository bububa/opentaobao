package tmallservice

import (
	"sync"
)

// TmallMallitemcenterServiceproductQueryResult 结构体
type TmallMallitemcenterServiceproductQueryResult struct {
	// 返回数据
	ResultDataList []ResultData `json:"result_data_list,omitempty" xml:"result_data_list>result_data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统是否异常
	SystemError bool `json:"system_error,omitempty" xml:"system_error,omitempty"`
	// 业务校验是否正常
	BusinessCheckFail bool `json:"business_check_fail,omitempty" xml:"business_check_fail,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallMallitemcenterServiceproductQueryResult = sync.Pool{
	New: func() any {
		return new(TmallMallitemcenterServiceproductQueryResult)
	},
}

// GetTmallMallitemcenterServiceproductQueryResult() 从对象池中获取TmallMallitemcenterServiceproductQueryResult
func GetTmallMallitemcenterServiceproductQueryResult() *TmallMallitemcenterServiceproductQueryResult {
	return poolTmallMallitemcenterServiceproductQueryResult.Get().(*TmallMallitemcenterServiceproductQueryResult)
}

// ReleaseTmallMallitemcenterServiceproductQueryResult 释放TmallMallitemcenterServiceproductQueryResult
func ReleaseTmallMallitemcenterServiceproductQueryResult(v *TmallMallitemcenterServiceproductQueryResult) {
	v.ResultDataList = v.ResultDataList[:0]
	v.Message = ""
	v.ErrorCode = ""
	v.SystemError = false
	v.BusinessCheckFail = false
	v.Success = false
	poolTmallMallitemcenterServiceproductQueryResult.Put(v)
}
