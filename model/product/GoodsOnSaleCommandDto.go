package product

// GoodsOnSaleCommandDto 结构体
type GoodsOnSaleCommandDto struct {
	// 批量上架商品id集合
	ExternalGoodsIdList []ExternalGoodsIdDto `json:"external_goods_id_list,omitempty" xml:"external_goods_id_list>external_goods_id_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}
