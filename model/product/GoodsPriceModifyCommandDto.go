package product

// GoodsPriceModifyCommandDto 结构体
type GoodsPriceModifyCommandDto struct {
	// 商品价格传输对象
	GoodsPriceList []GoodsPriceDto `json:"goods_price_list,omitempty" xml:"goods_price_list>goods_price_dto,omitempty"`
	// 外部批次ID，用于幂等
	ExternalBatchId string `json:"external_batch_id,omitempty" xml:"external_batch_id,omitempty"`
}
