package fenxiao

import (
	"sync"
)

// SubOrderDetail 结构体
type SubOrderDetail struct {
	// 建议废弃 Feature对象列表目前已有的属性： attr_key为 www，attr_value为1 表示是www子订单； attr_key为 wwwStoreCode，attr_value是www子订单发货的仓库编码； attr_key为 isWt，attr_value为1 表示是网厅子订单； attr_key为wtInfo,attr_value为网厅相关合约信息；  attr_key为 storeCode，attr_value为仓库信息；  attr_key为 erpHold，attr_value为1表示强管控中， attr_value为2表示分单完成；
	Features []FeatureDo `json:"features,omitempty" xml:"features>feature_do,omitempty"`
	// 商品商家编码，供应商发布产品是定义的产品编码
	ItemOuterId string `json:"item_outer_id,omitempty" xml:"item_outer_id,omitempty"`
	// SKU商家编码，供应商发布产品SKU是定义的编码
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// SKU属性值。如: 颜色:红色:别名;尺码:L:别名
	SkuProperties string `json:"sku_properties,omitempty" xml:"sku_properties,omitempty"`
	// 老的SKU属性值。如: 颜色:红色:别名;尺码:L:别名
	OldSkuProperties string `json:"old_sku_properties,omitempty" xml:"old_sku_properties,omitempty"`
	// 采购的产品标题。
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 产品的采购价格。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 废弃： 产品快照地址
	SnapshotUrl string `json:"snapshot_url,omitempty" xml:"snapshot_url,omitempty"`
	// 子单创建时间。格式 yyyy-MM-dd HH:mm:ss 。
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 采购单交易状态。 可选值： WAIT_BUYER_PAY(等待付款) WAIT_SELLER_CONFIRM(付款信息待确认，待发货) WAIT_SELLER_SEND_GOODS(已付款，待发货） WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)  TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_REFUNDING(退款中) TRADE_REFUNDED(已退款) PAID_FORBID_CONSIGN(已付款, 禁止发货(大快消行业)) WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持) PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功） PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 分销商应付金额。 计算公式：num(采购数量)*price(采购价) - 折扣。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 分销商实付金额。 total_fee（分销商应付金额）+ 改价。（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	DistributorPayment string `json:"distributor_payment,omitempty" xml:"distributor_payment,omitempty"`
	// 买家订单上对应的子单零售金额。 计算公式：消费者订单的成交价 * 购买数量 + 调整金额 - 优惠折扣（精确到2位小数;单位:元。如:200.07，表示:200元7分）
	BuyerPayment string `json:"buyer_payment,omitempty" xml:"buyer_payment,omitempty"`
	// 退款金额，单位元。（精确到2位小数;单位:元。如:12.23，表示:12元2角3分）
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 消费者的交易订单状态（代销采购单对应下游200订单状态）。 可选值： WAIT_SELLER_SEND_GOODS(已付款，待发货) WAIT_BUYER_CONFIRM_GOODS(已付款，已发货) TRADE_CLOSED(已退款成功) TRADE_REFUNDING(退款中) TRADE_FINISHED(交易成功) TRADE_CLOSED_BY_TAOBAO(交易关闭)
	Order200Status string `json:"order_200_status,omitempty" xml:"order_200_status,omitempty"`
	// 分销商店铺中宝贝一口价，消费者看见的宝贝价格。单位元 代销场景下存在，此价格不是实付款的价格。
	AuctionPrice string `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 发票应开金额。 子单的消费者实付分摊金额。根据买家实际付款去除邮费后，按各个子单(商品)金额比例进行分摊后的金额，仅供开发票时做票面金额参考。
	BillFee string `json:"bill_fee,omitempty" xml:"bill_fee,omitempty"`
	// 商品优惠类型：聚划算、秒杀或其他
	TcPreferentialType string `json:"tc_preferential_type,omitempty" xml:"tc_preferential_type,omitempty"`
	// 消费者订单优惠金额，始终为正数，单位是分，不带小数
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 采购子订单优惠活动类型。0=无优惠；1=限时折
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 消费者订单优惠金额，始终为正数，单位是元
	DiscountFeeYuan string `json:"discount_fee_yuan,omitempty" xml:"discount_fee_yuan,omitempty"`
	// sku属性值
	SkuPropertyVal string `json:"sku_property_val,omitempty" xml:"sku_property_val,omitempty"`
	// 分销商实付金额，单位：元
	DistributorPayPriceYuan string `json:"distributor_pay_price_yuan,omitempty" xml:"distributor_pay_price_yuan,omitempty"`
	// 退款金额，单位：元
	RefundFeeYuan string `json:"refund_fee_yuan,omitempty" xml:"refund_fee_yuan,omitempty"`
	// 确认收货金额，单位：元
	ConfirmPaidFeeYuan string `json:"confirm_paid_fee_yuan,omitempty" xml:"confirm_paid_fee_yuan,omitempty"`
	// 买家实付金额，单位：元
	BuyerPayPriceYuan string `json:"buyer_pay_price_yuan,omitempty" xml:"buyer_pay_price_yuan,omitempty"`
	// 老的sku属性值
	OldSkuPropertyVal string `json:"old_sku_property_val,omitempty" xml:"old_sku_property_val,omitempty"`
	// 支付宝交易号，担保交易使用。
	AlipayOrderNo string `json:"alipay_order_no,omitempty" xml:"alipay_order_no,omitempty"`
	// 图片地址
	SnapPath string `json:"snap_path,omitempty" xml:"snap_path,omitempty"`
	// 子采购单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 发票应开金额。 子单的消费者实付分摊金额。根据买家实际付款去除邮费后，按各个子单(商品)金额比例进行分摊后的金额，仅供开发票时做票面金额参考。单位：元
	BillFeeYuan string `json:"bill_fee_yuan,omitempty" xml:"bill_fee_yuan,omitempty"`
	// 单个商品价格，单位：元
	PriceYuan string `json:"price_yuan,omitempty" xml:"price_yuan,omitempty"`
	// 分销商应付金额，单位：元
	DistributorPriceYuan string `json:"distributor_price_yuan,omitempty" xml:"distributor_price_yuan,omitempty"`
	// 子采购单退款状态，枚举：RF_STATUS_NO_REFUND-没有申请退款；RF_STATUS_REFUNDING-退款中；RF_STATUS_END_REFUND-退款结束，退款成功或者退款关闭都会走到这里
	RefundStatusCode string `json:"refund_status_code,omitempty" xml:"refund_status_code,omitempty"`
	// 子采购单交易状态，枚举：AUDITING-意向单审核中；LOCKED-订单付款前锁定；WAIT_FOR_ALLOCATE_GOODS-待分货；ALLOCATING-分货中；WAIT_FOR_PAY-等待付款；PAID_WAIT_FOR_CONFIRM-已付款待确认付款操作；PAID-已确认付款，可直接发货；DELIVERED-已提交交货；CONSIGNED-已发货；RECEIVING-已提交确认收货，但收货未完成；ORDER_CLOSE-交易关闭；ORDER_SUCCESS-交易成功
	SubOrderStatusCode string `json:"sub_order_status_code,omitempty" xml:"sub_order_status_code,omitempty"`
	// 商品商家编码，供应商发布产品时定义的产品编码
	AuctionOutId string `json:"auction_out_id,omitempty" xml:"auction_out_id,omitempty"`
	// SKU商家编码，供应商发布产品SKU是定义的编码
	AuctionOutSkuId string `json:"auction_out_sku_id,omitempty" xml:"auction_out_sku_id,omitempty"`
	// 消费者订单子订单ID（经销模式下不存在该值）
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 子采购单编号
	FenxiaoId int64 `json:"fenxiao_id,omitempty" xml:"fenxiao_id,omitempty"`
	// 分销平台上的产品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 分销产品的SKU ID。 当存在时才会有值，建议使用sku_outer_id，sku_properties这两个值
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 消费者购买宝贝ID，不存在时为0。 2015年4月15日之前创建的采购单该字段都是0。
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 消费者购买宝贝SKU ID，不存在时为0。 2015年3月15日之前创建的采购单该字段都是0。
	AuctionSkuId int64 `json:"auction_sku_id,omitempty" xml:"auction_sku_id,omitempty"`
	// 产品的采购数量。取值范围:大于零的整数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 建议废弃：只有当担保交易支付时存在，为TP800的订单编号 子采购单id，淘宝交易主键，采购单未付款时为0.（只有支付宝 付款才有这个id，其余付款形式该字段为0）
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 优惠金额，始终为正数，单位是分，不带小数
	TcDiscountFee int64 `json:"tc_discount_fee,omitempty" xml:"tc_discount_fee,omitempty"`
	// 商品的卖出金额调整，金额增加时为正数，金额减少时为负数，单位是分，不带小数
	TcAdjustFee int64 `json:"tc_adjust_fee,omitempty" xml:"tc_adjust_fee,omitempty"`
	// 子订单留言信息
	TopMemo *TopMemoDto `json:"top_memo,omitempty" xml:"top_memo,omitempty"`
	// 买家实付金额，单位：分
	BuyerPayPrice int64 `json:"buyer_pay_price,omitempty" xml:"buyer_pay_price,omitempty"`
	// 购买数量
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
	// 前台子订单id
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
	// 确认收货金额，单位：分
	ConfirmPaidFee int64 `json:"confirm_paid_fee,omitempty" xml:"confirm_paid_fee,omitempty"`
	// 消费者购买宝贝ID，不存在时为0。 2015年4月15日之前创建的采购单该字段都是0。
	RealAuctionId int64 `json:"real_auction_id,omitempty" xml:"real_auction_id,omitempty"`
	// 消费者购买宝贝SKU ID，不存在时为0。 2015年3月15日之前创建的采购单该字段都是0。
	RealAuctionSkuId int64 `json:"real_auction_sku_id,omitempty" xml:"real_auction_sku_id,omitempty"`
	// 子采购单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 分销商应付金额，单位：分
	DistributorPrice int64 `json:"distributor_price,omitempty" xml:"distributor_price,omitempty"`
	// sku商家编码
	AuctionSellerSkuId int64 `json:"auction_seller_sku_id,omitempty" xml:"auction_seller_sku_id,omitempty"`
	// 分销商实付金额，单位：分
	DistributorPayPrice int64 `json:"distributor_pay_price,omitempty" xml:"distributor_pay_price,omitempty"`
	// 子采购单退款状态-数字形式，枚举：9-没有申请退款；20-退款中；21-退款结束，退款成功或者退款关闭都会走到这里
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 子采购单交易状态-数字表示，枚举：10-意向单审核中；20-订单付款前锁定；30-待分货；40-分货中；50-订单等待付款；55-已付款待确认付款操作；60-已付款；70-已提交交货；80-已发货；90-已提交确认收货，但收货未完成；100-交易关闭；110-交易成功
	SubOrderStatus int64 `json:"sub_order_status,omitempty" xml:"sub_order_status,omitempty"`
	// 二级供应商id
	SecondarySupplierId int64 `json:"secondary_supplier_id,omitempty" xml:"secondary_supplier_id,omitempty"`
	// tp单创单时间的货品采购单价
	TpCreateTimePrice float64 `json:"tp_create_time_price,omitempty" xml:"tp_create_time_price,omitempty"`
}

var poolSubOrderDetail = sync.Pool{
	New: func() any {
		return new(SubOrderDetail)
	},
}

// GetSubOrderDetail() 从对象池中获取SubOrderDetail
func GetSubOrderDetail() *SubOrderDetail {
	return poolSubOrderDetail.Get().(*SubOrderDetail)
}

// ReleaseSubOrderDetail 释放SubOrderDetail
func ReleaseSubOrderDetail(v *SubOrderDetail) {
	v.Features = v.Features[:0]
	v.ItemOuterId = ""
	v.SkuOuterId = ""
	v.SkuProperties = ""
	v.OldSkuProperties = ""
	v.Title = ""
	v.Price = ""
	v.SnapshotUrl = ""
	v.Created = ""
	v.Status = ""
	v.TotalFee = ""
	v.DistributorPayment = ""
	v.BuyerPayment = ""
	v.RefundFee = ""
	v.Order200Status = ""
	v.AuctionPrice = ""
	v.BillFee = ""
	v.TcPreferentialType = ""
	v.DiscountFee = ""
	v.PromotionType = ""
	v.DiscountFeeYuan = ""
	v.SkuPropertyVal = ""
	v.DistributorPayPriceYuan = ""
	v.RefundFeeYuan = ""
	v.ConfirmPaidFeeYuan = ""
	v.BuyerPayPriceYuan = ""
	v.OldSkuPropertyVal = ""
	v.AlipayOrderNo = ""
	v.SnapPath = ""
	v.GmtCreate = ""
	v.BillFeeYuan = ""
	v.PriceYuan = ""
	v.DistributorPriceYuan = ""
	v.RefundStatusCode = ""
	v.SubOrderStatusCode = ""
	v.AuctionOutId = ""
	v.AuctionOutSkuId = ""
	v.TcOrderId = 0
	v.FenxiaoId = 0
	v.ItemId = 0
	v.SkuId = 0
	v.AuctionId = 0
	v.AuctionSkuId = 0
	v.Num = 0
	v.Id = 0
	v.ScItemId = 0
	v.TcDiscountFee = 0
	v.TcAdjustFee = 0
	v.TopMemo = nil
	v.BuyerPayPrice = 0
	v.BuyNum = 0
	v.SubTcOrderId = 0
	v.ConfirmPaidFee = 0
	v.RealAuctionId = 0
	v.RealAuctionSkuId = 0
	v.SubOrderId = 0
	v.DistributorPrice = 0
	v.AuctionSellerSkuId = 0
	v.DistributorPayPrice = 0
	v.RefundStatus = 0
	v.SubOrderStatus = 0
	v.SecondarySupplierId = 0
	v.TpCreateTimePrice = 0
	poolSubOrderDetail.Put(v)
}
