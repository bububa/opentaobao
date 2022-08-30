package product

// ExternalGoodsIdDto 结构体
type ExternalGoodsIdDto struct {
	// 外部商品ID，用于标识外部系统每次提交过来的商品
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 交易猫商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
