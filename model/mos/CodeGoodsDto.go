package mos

// CodeGoodsDto 结构体
type CodeGoodsDto struct {
	// zi单号
	SubNo string `json:"sub_no,omitempty" xml:"sub_no,omitempty"`
	// 数量
	Quantity *BigDecimal `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
