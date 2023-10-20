package tbk

import (
	"sync"
)

// TaobaoTbkScVegasSendReportResult 结构体
type TaobaoTbkScVegasSendReportResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *RightsSendRptDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTbkScVegasSendReportResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScVegasSendReportResult)
	},
}

// GetTaobaoTbkScVegasSendReportResult() 从对象池中获取TaobaoTbkScVegasSendReportResult
func GetTaobaoTbkScVegasSendReportResult() *TaobaoTbkScVegasSendReportResult {
	return poolTaobaoTbkScVegasSendReportResult.Get().(*TaobaoTbkScVegasSendReportResult)
}

// ReleaseTaobaoTbkScVegasSendReportResult 释放TaobaoTbkScVegasSendReportResult
func ReleaseTaobaoTbkScVegasSendReportResult(v *TaobaoTbkScVegasSendReportResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolTaobaoTbkScVegasSendReportResult.Put(v)
}
