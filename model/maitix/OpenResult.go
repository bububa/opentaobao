package maitix

import (
	"sync"
)

// OpenResult 结构体
type OpenResult struct {
	// 结果
	ModelList []OpenExchangePointDto `json:"model_list,omitempty" xml:"model_list>open_exchange_point_dto,omitempty"`
	// 电子票对象
	ModelArrList []EticketDto `json:"model_arr_list,omitempty" xml:"model_arr_list>eticket_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 参数extMap
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 返回结果
	Model *DisEncrypt4CmbResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOpenResult = sync.Pool{
	New: func() any {
		return new(OpenResult)
	},
}

// GetOpenResult() 从对象池中获取OpenResult
func GetOpenResult() *OpenResult {
	return poolOpenResult.Get().(*OpenResult)
}

// ReleaseOpenResult 释放OpenResult
func ReleaseOpenResult(v *OpenResult) {
	v.ModelList = v.ModelList[:0]
	v.ModelArrList = v.ModelArrList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ExtMap = ""
	v.Model = nil
	v.Success = false
	poolOpenResult.Put(v)
}
