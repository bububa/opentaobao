package tmallchannel

// PageResultDto 结构体
type PageResultDto struct {
	// 产品信息
	ProductList []ProductTopDto `json:"product_list,omitempty" xml:"product_list>product_top_dto,omitempty"`
	// 异常信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 是否查询成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
