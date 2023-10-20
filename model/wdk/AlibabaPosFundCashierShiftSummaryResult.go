package wdk

import (
	"sync"
)

// AlibabaPosFundCashierShiftSummaryResult 结构体
type AlibabaPosFundCashierShiftSummaryResult struct {
	// 模型
	Model []CashierShiftFundSummaryDto `json:"model,omitempty" xml:"model>cashier_shift_fund_summary_dto,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 扩展字段
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaPosFundCashierShiftSummaryResult = sync.Pool{
	New: func() any {
		return new(AlibabaPosFundCashierShiftSummaryResult)
	},
}

// GetAlibabaPosFundCashierShiftSummaryResult() 从对象池中获取AlibabaPosFundCashierShiftSummaryResult
func GetAlibabaPosFundCashierShiftSummaryResult() *AlibabaPosFundCashierShiftSummaryResult {
	return poolAlibabaPosFundCashierShiftSummaryResult.Get().(*AlibabaPosFundCashierShiftSummaryResult)
}

// ReleaseAlibabaPosFundCashierShiftSummaryResult 释放AlibabaPosFundCashierShiftSummaryResult
func ReleaseAlibabaPosFundCashierShiftSummaryResult(v *AlibabaPosFundCashierShiftSummaryResult) {
	v.Model = v.Model[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.BizExtMap = ""
	v.Success = false
	poolAlibabaPosFundCashierShiftSummaryResult.Put(v)
}
