package lstfundbill

// LstFundBillOrderDto 结构体
type LstFundBillOrderDto struct {
	// 主订单id
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 结算总金额(含佣)
	TotalSettleAmtHy string `json:"total_settle_amt_hy,omitempty" xml:"total_settle_amt_hy,omitempty"`
	// 结算总金额(扣佣)
	TotalSettleAmt string `json:"total_settle_amt,omitempty" xml:"total_settle_amt,omitempty"`
	// 销售总金额
	TotalItemAmt string `json:"total_item_amt,omitempty" xml:"total_item_amt,omitempty"`
	// 运费
	Carriage string `json:"carriage,omitempty" xml:"carriage,omitempty"`
	// 退款运费
	RefundCarriage string `json:"refund_carriage,omitempty" xml:"refund_carriage,omitempty"`
	// 交易方式
	DealType string `json:"deal_type,omitempty" xml:"deal_type,omitempty"`
	// 支付宝入账金额
	AlipayFee string `json:"alipay_fee,omitempty" xml:"alipay_fee,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单发货时间
	GmtShipped string `json:"gmt_shipped,omitempty" xml:"gmt_shipped,omitempty"`
	// 确认收货时间
	GmtConfirmed string `json:"gmt_confirmed,omitempty" xml:"gmt_confirmed,omitempty"`
	// 订单完成时间
	GmtCompleted string `json:"gmt_completed,omitempty" xml:"gmt_completed,omitempty"`
	// 订单结算时间
	GmtSettled string `json:"gmt_settled,omitempty" xml:"gmt_settled,omitempty"`
	// 子订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 产品名称
	CspuName string `json:"cspu_name,omitempty" xml:"cspu_name,omitempty"`
	// cspuId
	CspuId string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
	// 规格
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// 条形码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 仓库类型
	WarehouseType string `json:"warehouse_type,omitempty" xml:"warehouse_type,omitempty"`
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 结算金额
	SettleAmt string `json:"settle_amt,omitempty" xml:"settle_amt,omitempty"`
	// 销售单价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 销售金额
	ItemAmt string `json:"item_amt,omitempty" xml:"item_amt,omitempty"`
	// 商品佣金
	ItemCommissionAmt string `json:"item_commission_amt,omitempty" xml:"item_commission_amt,omitempty"`
	// 物流佣金
	CarriageCommissionAmt string `json:"carriage_commission_amt,omitempty" xml:"carriage_commission_amt,omitempty"`
	// 无忧购佣金
	WygCommissionAmt string `json:"wyg_commission_amt,omitempty" xml:"wyg_commission_amt,omitempty"`
	// 金融技术服务佣金
	JrfwCommissionAmt string `json:"jrfw_commission_amt,omitempty" xml:"jrfw_commission_amt,omitempty"`
	// 退款单号
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 订单退款状态
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 货款退款金额
	RefundAmt string `json:"refund_amt,omitempty" xml:"refund_amt,omitempty"`
	// 优惠金额
	DiscountAmt string `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
	// 赠品换购
	MkZphg string `json:"mk_zphg,omitempty" xml:"mk_zphg,omitempty"`
	// 店铺满减
	MkShopMj string `json:"mk_shop_mj,omitempty" xml:"mk_shop_mj,omitempty"`
	// 店铺包邮
	MkShopBy string `json:"mk_shop_by,omitempty" xml:"mk_shop_by,omitempty"`
	// 单品优惠
	MkDpyh string `json:"mk_dpyh,omitempty" xml:"mk_dpyh,omitempty"`
	// 限时促销
	MkXscx string `json:"mk_xscx,omitempty" xml:"mk_xscx,omitempty"`
	// 采购券
	MkPurchaseCoupon string `json:"mk_purchase_coupon,omitempty" xml:"mk_purchase_coupon,omitempty"`
	// 品牌券
	MkBrandCoupon string `json:"mk_brand_coupon,omitempty" xml:"mk_brand_coupon,omitempty"`
	// 跨店满减
	MkKdmj string `json:"mk_kdmj,omitempty" xml:"mk_kdmj,omitempty"`
	// 跨店满赠
	MkKdmz string `json:"mk_kdmz,omitempty" xml:"mk_kdmz,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 小店名称
	LeadsName string `json:"leads_name,omitempty" xml:"leads_name,omitempty"`
	// 小店id
	LeadsId string `json:"leads_id,omitempty" xml:"leads_id,omitempty"`
	// 省
	LogisticsProv string `json:"logistics_prov,omitempty" xml:"logistics_prov,omitempty"`
	// 市
	LogisticsCity string `json:"logistics_city,omitempty" xml:"logistics_city,omitempty"`
	// 区
	LogisticsArea string `json:"logistics_area,omitempty" xml:"logistics_area,omitempty"`
	// 是否自有商品
	SelfOffer string `json:"self_offer,omitempty" xml:"self_offer,omitempty"`
	// 配送商信息
	DeliveryName string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 商家名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 商品佣金费率
	ItemCommissionRate string `json:"item_commission_rate,omitempty" xml:"item_commission_rate,omitempty"`
	// 物流佣金费率
	PostCommissionRate string `json:"post_commission_rate,omitempty" xml:"post_commission_rate,omitempty"`
	// 商品佣金退佣金额
	ItemCommissionRefundAmt string `json:"item_commission_refund_amt,omitempty" xml:"item_commission_refund_amt,omitempty"`
	// 零售通商品的唯一识别id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 销售数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 退货数量
	RefundQuantity int64 `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
}
