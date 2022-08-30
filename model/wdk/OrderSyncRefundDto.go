package wdk

// OrderSyncRefundDto 结构体
type OrderSyncRefundDto struct {
	// 退款渠道列表
	RefundChannelList []OrderSyncRefundChannel `json:"refund_channel_list,omitempty" xml:"refund_channel_list>order_sync_refund_channel,omitempty"`
	// 库存单位购买数量
	BuyAmountStock string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
	// 商品skucode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 非标品购买数量
	NsQuantity string `json:"ns_quantity,omitempty" xml:"ns_quantity,omitempty"`
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 库存单位拣货数量
	PickAmountStock string `json:"pick_amount_stock,omitempty" xml:"pick_amount_stock,omitempty"`
	// 退货数量
	RefundQuantity string `json:"refund_quantity,omitempty" xml:"refund_quantity,omitempty"`
	// 退款时间
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// 非标品单位
	SellUnit string `json:"sell_unit,omitempty" xml:"sell_unit,omitempty"`
	// 店铺码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 同步状态
	SyncStatus string `json:"sync_status,omitempty" xml:"sync_status,omitempty"`
	// 跑批时间
	BatchTime string `json:"batch_time,omitempty" xml:"batch_time,omitempty"`
	// 实际取货数量
	ActualPickQuantity string `json:"actual_pick_quantity,omitempty" xml:"actual_pick_quantity,omitempty"`
	// 实际退货数量
	ActualRefundQuantity string `json:"actual_refund_quantity,omitempty" xml:"actual_refund_quantity,omitempty"`
	// 履约状态
	PromiseStatus string `json:"promise_status,omitempty" xml:"promise_status,omitempty"`
	// stockUnit
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// saleUnit
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 附加属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 交易属性
	TradeAttributes string `json:"trade_attributes,omitempty" xml:"trade_attributes,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 外部子订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部主订单号
	OutMianOrderId string `json:"out_mian_order_id,omitempty" xml:"out_mian_order_id,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部skuCode
	OutSkuCode string `json:"out_sku_code,omitempty" xml:"out_sku_code,omitempty"`
	// 外部退款id
	OutRefundId string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 退款业务单号
	BizMainRefundId string `json:"biz_main_refund_id,omitempty" xml:"biz_main_refund_id,omitempty"`
	// 外部退款单号
	OutMainRefundId string `json:"out_main_refund_id,omitempty" xml:"out_main_refund_id,omitempty"`
	// 退单状态 APPLY：申请，待审核，REFUSE：审核拒绝，AGREE：审核同意（退款中），SUCCESS：退款成功，CANCEL：取消申请，PLATFROM_DEAL：平台仲裁中，FAILED：退款失败
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 扩展属性map
	RefundAttributes string `json:"refund_attributes,omitempty" xml:"refund_attributes,omitempty"`
	// 供应商code
	SourceMerchantCode string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
	// 盒马子订单号
	OriginOrderId int64 `json:"origin_order_id,omitempty" xml:"origin_order_id,omitempty"`
	// 盒马主订单号
	OriginParentId int64 `json:"origin_parent_id,omitempty" xml:"origin_parent_id,omitempty"`
	// 原购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 退款单id
	RefundOrderId int64 `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 退款类型
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// 淘系子订单号
	TbBizOrderId int64 `json:"tb_biz_order_id,omitempty" xml:"tb_biz_order_id,omitempty"`
	// 退款运费，单位为分
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
	// 淘系主订单号
	TbBizParentId int64 `json:"tb_biz_parent_id,omitempty" xml:"tb_biz_parent_id,omitempty"`
	// 2:app 3:pos
	OrderClient int64 `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 渠道来源 3：饿了么 4：盒马
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 纠纷类型 1：售中  2：售后
	DisputeType int64 `json:"dispute_type,omitempty" xml:"dispute_type,omitempty"`
	// 退款销售数量
	RefundSaleQuantity int64 `json:"refund_sale_quantity,omitempty" xml:"refund_sale_quantity,omitempty"`
}
