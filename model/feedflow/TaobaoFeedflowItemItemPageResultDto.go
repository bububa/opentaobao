package feedflow

// TaobaoFeedflowItemItemPageResultDto 结构体
type TaobaoFeedflowItemItemPageResultDto struct {
	// 商品列表
	ItemList []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 商品总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
