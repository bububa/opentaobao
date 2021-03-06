package logistic

// GoodsItem 结构体
type GoodsItem struct {
	// 货品ID
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 子订单编号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 货品数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
