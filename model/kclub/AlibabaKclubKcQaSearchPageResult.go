package kclub

// AlibabakclubkcqasearchpageResult 结构体
type AlibabakclubkcqasearchpageResult struct {
	// 错误编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页数据
	Data *Paging `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
