package alihealth2

// TopGoodsDto 结构体
type TopGoodsDto struct {
	// 商品名
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 商品编码
	GoodsCode string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
	// 商品规格
	GoodsSpec string `json:"goods_spec,omitempty" xml:"goods_spec,omitempty"`
	// 包装
	PackUnitName string `json:"pack_unit_name,omitempty" xml:"pack_unit_name,omitempty"`
	// 包装数
	PackQty int64 `json:"pack_qty,omitempty" xml:"pack_qty,omitempty"`
	// 批号
	LotNo string `json:"lot_no,omitempty" xml:"lot_no,omitempty"`
	// 生产日期
	ProductDate string `json:"product_date,omitempty" xml:"product_date,omitempty"`
	// 有效期
	ValidityDate string `json:"validity_date,omitempty" xml:"validity_date,omitempty"`
	// 收到数量
	ReceiveQty int64 `json:"receive_qty,omitempty" xml:"receive_qty,omitempty"`
	// 单价
	UnitPriceStr string `json:"unit_price_str,omitempty" xml:"unit_price_str,omitempty"`
	// 金额
	AmountMoney int64 `json:"amount_money,omitempty" xml:"amount_money,omitempty"`
}
