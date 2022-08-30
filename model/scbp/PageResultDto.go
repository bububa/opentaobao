package scbp

// PageResultDto 结构体
type PageResultDto struct {
	// 返回数据
	ResultList []AdProductDto `json:"result_list,omitempty" xml:"result_list>ad_product_dto,omitempty"`
	// 错误信息
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
