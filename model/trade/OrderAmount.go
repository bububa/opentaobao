package trade

// OrderAmount 结构体
type OrderAmount struct {
	// 子交易订单编号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// SKU的值。如：机身颜色:黑色;手机套餐:官方标配
	SkuPropertiesName string `json:"sku_properties_name,omitempty" xml:"sku_properties_name,omitempty"`
	// 子交易订单中购买商品的数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 子订单的系统优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 卖家手工调整的子订单的优惠金额.格式为:1.01;单位:元;精确到小数点后两位.
	AdjustFee string `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
	// 子订单的系统优惠的名称，对应于discount_fee的名称
	PromotionName string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// 子订单对应的商品数字id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 子订单对应的商品的sku_id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。计算公式如下：payment = price * num + adjust_fee - discount_fee + post_fee(邮费，单笔子订单时子订单实付金额包含邮费，多笔子订单时不包含邮费)；对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 分摊之后的实付金额
	DivideOrderFee string `json:"divide_order_fee,omitempty" xml:"divide_order_fee,omitempty"`
	// 优惠分摊
	PartMjzDiscount string `json:"part_mjz_discount,omitempty" xml:"part_mjz_discount,omitempty"`
}
