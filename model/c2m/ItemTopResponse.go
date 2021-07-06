package c2m

// ItemTopResponse 结构体
type ItemTopResponse struct {
	// 返回的商品信息
	ItemTopDTOList []ItemTopDto `json:"item_top_d_t_o_list,omitempty" xml:"item_top_d_t_o_list>item_top_dto,omitempty"`
	// 总共多少条
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 每页多少条
	PagSize int64 `json:"pag_size,omitempty" xml:"pag_size,omitempty"`
}
