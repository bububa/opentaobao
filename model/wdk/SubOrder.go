package wdk

import (
	"sync"
)

// SubOrder 结构体
type SubOrder struct {
	// 营销优惠明细
	DiscountInfos []DiscountInfo `json:"discount_infos,omitempty" xml:"discount_infos>discount_info,omitempty"`
	// 外部子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 处理方式
	HandlingType string `json:"handling_type,omitempty" xml:"handling_type,omitempty"`
	// 盒马子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 促销信息(json格式)。例如:[{&#34;activity_id&#34;:&#34;1234&#34;,&#34;activity_name&#34;:&#34;五一促销&#34;,&#34;activity_type&#34;:1,&#34;activity_desc&#34;:&#34;优惠卡券&#34;}]
	PromotionInfo string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 非标品销售单位
	SellUnit string `json:"sell_unit,omitempty" xml:"sell_unit,omitempty"`
	// 非标品购买数量
	NsQuantity string `json:"ns_quantity,omitempty" xml:"ns_quantity,omitempty"`
	// 库存单位拣货数量
	PickAmountStock string `json:"pick_amount_stock,omitempty" xml:"pick_amount_stock,omitempty"`
	// 库存单位购买数量
	BuyAmountStock string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
	// memberPoint
	MemberPoint string `json:"member_point,omitempty" xml:"member_point,omitempty"`
	// 子订单类型，当前取值[COMMON|GIFT],分别表示 普通|赠品 订单
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 揽件
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// statusChangeTime
	StatusChangeTime string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// tradeSubAttributes
	TradeSubAttributes string `json:"trade_sub_attributes,omitempty" xml:"trade_sub_attributes,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部skucode
	OutSkuCode string `json:"out_sku_code,omitempty" xml:"out_sku_code,omitempty"`
	// 销售商品数量
	SaleQuantity int64 `json:"sale_quantity,omitempty" xml:"sale_quantity,omitempty"`
	// 销售单价
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 实际支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 原始金额
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 子单优惠金额商家分摊
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 子单优惠金额平台分摊
	PlatformDiscountFee int64 `json:"platform_discount_fee,omitempty" xml:"platform_discount_fee,omitempty"`
	// 子单商品总重量
	TotalWeight int64 `json:"total_weight,omitempty" xml:"total_weight,omitempty"`
	// 业务子订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 促销优惠总金额
	PromotionDiscountAmt int64 `json:"promotion_discount_amt,omitempty" xml:"promotion_discount_amt,omitempty"`
	// 商品id
	ItemCode int64 `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 正向：售价金额（购买数量*原售价）。逆向：退款金额
	OriginalAmt int64 `json:"original_amt,omitempty" xml:"original_amt,omitempty"`
	// 父订单id
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 商品原价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 销售类型（正向销售1：逆向销售2。本接口应返回2）
	TrdType int64 `json:"trd_type,omitempty" xml:"trd_type,omitempty"`
	// 会员价优惠金额
	MemberDiscountAmt int64 `json:"member_discount_amt,omitempty" xml:"member_discount_amt,omitempty"`
	// 其它分摊优惠金额
	ShareDiscountAmt int64 `json:"share_discount_amt,omitempty" xml:"share_discount_amt,omitempty"`
	// 淘宝子订单号
	TbBizOrderId int64 `json:"tb_biz_order_id,omitempty" xml:"tb_biz_order_id,omitempty"`
	// 淘鲜达平台优惠券中平台出资金额，单位分
	TxdPmtAmt int64 `json:"txd_pmt_amt,omitempty" xml:"txd_pmt_amt,omitempty"`
	// 拣货金额
	PickAmt int64 `json:"pick_amt,omitempty" xml:"pick_amt,omitempty"`
}

var poolSubOrder = sync.Pool{
	New: func() any {
		return new(SubOrder)
	},
}

// GetSubOrder() 从对象池中获取SubOrder
func GetSubOrder() *SubOrder {
	return poolSubOrder.Get().(*SubOrder)
}

// ReleaseSubOrder 释放SubOrder
func ReleaseSubOrder(v *SubOrder) {
	v.DiscountInfos = v.DiscountInfos[:0]
	v.SubOutOrderId = ""
	v.SkuCode = ""
	v.HandlingType = ""
	v.SubBizOrderId = ""
	v.PromotionInfo = ""
	v.StoreId = ""
	v.SellUnit = ""
	v.NsQuantity = ""
	v.PickAmountStock = ""
	v.BuyAmountStock = ""
	v.MemberPoint = ""
	v.OrderType = ""
	v.OrderStatus = ""
	v.StatusChangeTime = ""
	v.StockUnit = ""
	v.SaleUnit = ""
	v.TradeSubAttributes = ""
	v.OutOrderId = ""
	v.OutSkuCode = ""
	v.SaleQuantity = 0
	v.SalePrice = 0
	v.PayFee = 0
	v.OriginFee = 0
	v.DiscountFee = 0
	v.MerchantDiscountFee = 0
	v.PlatformDiscountFee = 0
	v.TotalWeight = 0
	v.BizOrderId = 0
	v.PromotionDiscountAmt = 0
	v.ItemCode = 0
	v.OriginalAmt = 0
	v.ParentId = 0
	v.Price = 0
	v.Quantity = 0
	v.TrdType = 0
	v.MemberDiscountAmt = 0
	v.ShareDiscountAmt = 0
	v.TbBizOrderId = 0
	v.TxdPmtAmt = 0
	v.PickAmt = 0
	poolSubOrder.Put(v)
}
