package foodscan

// AlibabaFootscanMiniReportFragmentSecondMtopResult 
type AlibabaFootscanMiniReportFragmentSecondMtopResult struct {
    // 成功
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 返回i数据
    Data   *AlibabaFootscanMiniReportFragmentSecondData `json:"data,omitempty" xml:"data,omitempty"`
    // 成功
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}