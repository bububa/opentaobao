package foodscan

// AlibabaFootscanMiniReportListMtopResult 结构体
type AlibabaFootscanMiniReportListMtopResult struct {
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据
	Data *AlibabaFootscanMiniReportListData `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
