package trade

// RefundGoodsSubOrderDetail 结构体
type RefundGoodsSubOrderDetail struct {
	// 子订单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 履约单号
	FulfillId string `json:"fulfill_id,omitempty" xml:"fulfill_id,omitempty"`
	// 退货单号
	RefundGoodsId string `json:"refund_goods_id,omitempty" xml:"refund_goods_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 期望取货开始时间
	ExpectFulfilStartTime string `json:"expect_fulfil_start_time,omitempty" xml:"expect_fulfil_start_time,omitempty"`
	// 期望取货结束时间
	ExpectFulfilEndTime string `json:"expect_fulfil_end_time,omitempty" xml:"expect_fulfil_end_time,omitempty"`
	// 期望退货件数
	ExpectRefundAmount string `json:"expect_refund_amount,omitempty" xml:"expect_refund_amount,omitempty"`
	// 实际退货件数
	ActualRefundAmount string `json:"actual_refund_amount,omitempty" xml:"actual_refund_amount,omitempty"`
	// 履约取货件数
	FulfilRefundAmount string `json:"fulfil_refund_amount,omitempty" xml:"fulfil_refund_amount,omitempty"`
	// 同意退货件数
	AgreeRefundAmount string `json:"agree_refund_amount,omitempty" xml:"agree_refund_amount,omitempty"`
	// 退货商品skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品名称
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 商品主计价单位
	ItemIu string `json:"item_iu,omitempty" xml:"item_iu,omitempty"`
	// 商品销售单位
	ItemBu string `json:"item_bu,omitempty" xml:"item_bu,omitempty"`
	// 退货类型(1:"闪退";2: "标准")
	RefundSpeedType int64 `json:"refund_speed_type,omitempty" xml:"refund_speed_type,omitempty"`
	// 是否称重商品 1：是，0:否
	WeightItem int64 `json:"weight_item,omitempty" xml:"weight_item,omitempty"`
	// Create(10, "已发起退货"),SellerAgree(20, "卖家已同意"),OperatorAccept(30, "配送员已接单"),OperatorReceive(40, "配送员已取货"),Finish(50, "退货结束");
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// Init(0,"初始状态"),NoTakeGoods(1, "无需取货"),FailTakeGoods(2, "取货失败"),Timeout(3, "超时关闭"),EnterDock(4, "已入站");
	FinishType int64 `json:"finish_type,omitempty" xml:"finish_type,omitempty"`
}
