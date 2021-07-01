package aecreatives

// AliexpressAffiliateImageSearchResponse 结构体
type AliexpressAffiliateImageSearchResponse struct {
	// 返回结果状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 图搜结果
	Data *TrafficImageSearchResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 默认描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
