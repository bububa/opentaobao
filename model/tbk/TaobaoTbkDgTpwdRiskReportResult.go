package tbk

import (
	"sync"
)

// TaobaoTbkDgTpwdRiskReportResult 结构体
type TaobaoTbkDgTpwdRiskReportResult struct {
	// x
	StatusList []DataMap `json:"status_list,omitempty" xml:"status_list>data_map,omitempty"`
}

var poolTaobaoTbkDgTpwdRiskReportResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgTpwdRiskReportResult)
	},
}

// GetTaobaoTbkDgTpwdRiskReportResult() 从对象池中获取TaobaoTbkDgTpwdRiskReportResult
func GetTaobaoTbkDgTpwdRiskReportResult() *TaobaoTbkDgTpwdRiskReportResult {
	return poolTaobaoTbkDgTpwdRiskReportResult.Get().(*TaobaoTbkDgTpwdRiskReportResult)
}

// ReleaseTaobaoTbkDgTpwdRiskReportResult 释放TaobaoTbkDgTpwdRiskReportResult
func ReleaseTaobaoTbkDgTpwdRiskReportResult(v *TaobaoTbkDgTpwdRiskReportResult) {
	v.StatusList = v.StatusList[:0]
	poolTaobaoTbkDgTpwdRiskReportResult.Put(v)
}
