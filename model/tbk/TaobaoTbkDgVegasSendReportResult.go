package tbk

import (
	"sync"
)

// TaobaoTbkDgVegasSendReportResult 结构体
type TaobaoTbkDgVegasSendReportResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *RightsSendRptDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTbkDgVegasSendReportResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasSendReportResult)
	},
}

// GetTaobaoTbkDgVegasSendReportResult() 从对象池中获取TaobaoTbkDgVegasSendReportResult
func GetTaobaoTbkDgVegasSendReportResult() *TaobaoTbkDgVegasSendReportResult {
	return poolTaobaoTbkDgVegasSendReportResult.Get().(*TaobaoTbkDgVegasSendReportResult)
}

// ReleaseTaobaoTbkDgVegasSendReportResult 释放TaobaoTbkDgVegasSendReportResult
func ReleaseTaobaoTbkDgVegasSendReportResult(v *TaobaoTbkDgVegasSendReportResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolTaobaoTbkDgVegasSendReportResult.Put(v)
}
