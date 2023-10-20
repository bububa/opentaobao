package tbk

import (
	"sync"
)

// TaobaoTbkScUcrowdReportGetRpcResult 结构体
type TaobaoTbkScUcrowdReportGetRpcResult struct {
	// ucrowd_report
	UcrowdReport *DwsUnionAppDto `json:"ucrowd_report,omitempty" xml:"ucrowd_report,omitempty"`
}

var poolTaobaoTbkScUcrowdReportGetRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdReportGetRpcResult)
	},
}

// GetTaobaoTbkScUcrowdReportGetRpcResult() 从对象池中获取TaobaoTbkScUcrowdReportGetRpcResult
func GetTaobaoTbkScUcrowdReportGetRpcResult() *TaobaoTbkScUcrowdReportGetRpcResult {
	return poolTaobaoTbkScUcrowdReportGetRpcResult.Get().(*TaobaoTbkScUcrowdReportGetRpcResult)
}

// ReleaseTaobaoTbkScUcrowdReportGetRpcResult 释放TaobaoTbkScUcrowdReportGetRpcResult
func ReleaseTaobaoTbkScUcrowdReportGetRpcResult(v *TaobaoTbkScUcrowdReportGetRpcResult) {
	v.UcrowdReport = nil
	poolTaobaoTbkScUcrowdReportGetRpcResult.Put(v)
}
