package wlbimports

import (
	"sync"
)

// TaobaoWlbImportsVasIdentityResultTopResult 结构体
type TaobaoWlbImportsVasIdentityResultTopResult struct {
	// 返回错误信息
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 子错误编码
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 鉴定结果
	Result *TaobaoWlbImportsVasIdentityResultResult `json:"result,omitempty" xml:"result,omitempty"`
	// 总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWlbImportsVasIdentityResultTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsVasIdentityResultTopResult)
	},
}

// GetTaobaoWlbImportsVasIdentityResultTopResult() 从对象池中获取TaobaoWlbImportsVasIdentityResultTopResult
func GetTaobaoWlbImportsVasIdentityResultTopResult() *TaobaoWlbImportsVasIdentityResultTopResult {
	return poolTaobaoWlbImportsVasIdentityResultTopResult.Get().(*TaobaoWlbImportsVasIdentityResultTopResult)
}

// ReleaseTaobaoWlbImportsVasIdentityResultTopResult 释放TaobaoWlbImportsVasIdentityResultTopResult
func ReleaseTaobaoWlbImportsVasIdentityResultTopResult(v *TaobaoWlbImportsVasIdentityResultTopResult) {
	v.SubErrorMsg = ""
	v.ErrorCode = ""
	v.SubErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	v.TotalResults = 0
	v.Success = false
	poolTaobaoWlbImportsVasIdentityResultTopResult.Put(v)
}
