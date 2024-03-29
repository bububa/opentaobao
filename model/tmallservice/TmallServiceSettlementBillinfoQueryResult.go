package tmallservice

import (
	"sync"
)

// TmallServiceSettlementBillinfoQueryResult 结构体
type TmallServiceSettlementBillinfoQueryResult struct {
	// 错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误提示
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 工单结算信息
	Value *WorkcardBillInfoDto `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServiceSettlementBillinfoQueryResult = sync.Pool{
	New: func() any {
		return new(TmallServiceSettlementBillinfoQueryResult)
	},
}

// GetTmallServiceSettlementBillinfoQueryResult() 从对象池中获取TmallServiceSettlementBillinfoQueryResult
func GetTmallServiceSettlementBillinfoQueryResult() *TmallServiceSettlementBillinfoQueryResult {
	return poolTmallServiceSettlementBillinfoQueryResult.Get().(*TmallServiceSettlementBillinfoQueryResult)
}

// ReleaseTmallServiceSettlementBillinfoQueryResult 释放TmallServiceSettlementBillinfoQueryResult
func ReleaseTmallServiceSettlementBillinfoQueryResult(v *TmallServiceSettlementBillinfoQueryResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Value = nil
	v.Success = false
	poolTmallServiceSettlementBillinfoQueryResult.Put(v)
}
