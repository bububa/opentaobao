package foodscan

// AlibabaFootscanMiniSavedMtopResult 结构体
type AlibabaFootscanMiniSavedMtopResult struct {
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否更新成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
