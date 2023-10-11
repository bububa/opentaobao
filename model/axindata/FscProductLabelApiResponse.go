package axindata

// FscProductLabelApiResponse 结构体
type FscProductLabelApiResponse struct {
	// 返回数据
	Data []FscProductLabelApiDto `json:"data,omitempty" xml:"data>fsc_product_label_api_dto,omitempty"`
	// 数据总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
