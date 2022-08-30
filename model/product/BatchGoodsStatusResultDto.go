package product

// BatchGoodsStatusResultDto 结构体
type BatchGoodsStatusResultDto struct {
	// 商品状态列表
	GoodsStatusList []GoodsStatusDto `json:"goods_status_list,omitempty" xml:"goods_status_list>goods_status_dto,omitempty"`
}
