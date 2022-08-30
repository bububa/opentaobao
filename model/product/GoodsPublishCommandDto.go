package product

// GoodsPublishCommandDto 结构体
type GoodsPublishCommandDto struct {
	// 商品发布数据体
	GoodsList []GoodsPublishDto `json:"goods_list,omitempty" xml:"goods_list>goods_publish_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}
