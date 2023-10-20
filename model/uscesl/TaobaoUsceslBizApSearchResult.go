package uscesl

// TaobaousceslbizapsearchResult 结构体
type TaobaousceslbizapsearchResult struct {
	// 返回对象list
	TargetList []Target `json:"target_list,omitempty" xml:"target_list>target,omitempty"`
	// 调用返回编码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误编码
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数
	TotalPages int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// 每页显示数
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 本次调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
