package alihouse

import (
	"sync"
)

// TmallAlihouseTradeCouponOrderContractCredentialsQueryResult 结构体
type TmallAlihouseTradeCouponOrderContractCredentialsQueryResult struct {
	// 结果信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 凭证对象
	Data *StsCredentialsDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallAlihouseTradeCouponOrderContractCredentialsQueryResult = sync.Pool{
	New: func() any {
		return new(TmallAlihouseTradeCouponOrderContractCredentialsQueryResult)
	},
}

// GetTmallAlihouseTradeCouponOrderContractCredentialsQueryResult() 从对象池中获取TmallAlihouseTradeCouponOrderContractCredentialsQueryResult
func GetTmallAlihouseTradeCouponOrderContractCredentialsQueryResult() *TmallAlihouseTradeCouponOrderContractCredentialsQueryResult {
	return poolTmallAlihouseTradeCouponOrderContractCredentialsQueryResult.Get().(*TmallAlihouseTradeCouponOrderContractCredentialsQueryResult)
}

// ReleaseTmallAlihouseTradeCouponOrderContractCredentialsQueryResult 释放TmallAlihouseTradeCouponOrderContractCredentialsQueryResult
func ReleaseTmallAlihouseTradeCouponOrderContractCredentialsQueryResult(v *TmallAlihouseTradeCouponOrderContractCredentialsQueryResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolTmallAlihouseTradeCouponOrderContractCredentialsQueryResult.Put(v)
}
