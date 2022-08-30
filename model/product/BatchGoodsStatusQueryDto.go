package product

// BatchGoodsStatusQueryDto 结构体
type BatchGoodsStatusQueryDto struct {
	// 商品ID数组，size最大支持100
	GoodsIdList []int64 `json:"goods_id_list,omitempty" xml:"goods_id_list>int64,omitempty"`
	// 每页数量，支持最大100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页数
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}
