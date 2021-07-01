package foodscan

// AlibabaFootscanMiniImageUploadMtopResult 结构体
type AlibabaFootscanMiniImageUploadMtopResult struct {
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回i数据
	Data *AlibabaFootscanMiniImageUploadData `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
