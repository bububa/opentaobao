package tmallnr

import (
	"sync"
)

// NrResult 结构体
type NrResult struct {
	// 系统自动生成
	ResultDatas []ResultData `json:"result_datas,omitempty" xml:"result_datas>result_data,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode2 string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 取消是否成功
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// isSuccess
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolNrResult = sync.Pool{
	New: func() any {
		return new(NrResult)
	},
}

// GetNrResult() 从对象池中获取NrResult
func GetNrResult() *NrResult {
	return poolNrResult.Get().(*NrResult)
}

// ReleaseNrResult 释放NrResult
func ReleaseNrResult(v *NrResult) {
	v.ResultDatas = v.ResultDatas[:0]
	v.ErrorMessage = ""
	v.ErrorCode2 = ""
	v.ErrorCode = ""
	v.ResultData = false
	v.Success = false
	v.IsSuccess = false
	v.Succ = false
	poolNrResult.Put(v)
}
