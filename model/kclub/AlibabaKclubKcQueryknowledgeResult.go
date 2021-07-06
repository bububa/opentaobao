package kclub

// AlibabaKclubKcQueryknowledgeResult 结构体
type AlibabaKclubKcQueryknowledgeResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数据
	Data *Paging `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
