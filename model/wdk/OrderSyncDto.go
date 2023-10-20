package wdk

import (
	"sync"
)

// OrderSyncDto 结构体
type OrderSyncDto struct {
	// subOrders
	SubOrders []SubOrderSyncDto `json:"sub_orders,omitempty" xml:"sub_orders>sub_order_sync_dto,omitempty"`
	// payChannels
	PayChannels []PayChannel `json:"pay_channels,omitempty" xml:"pay_channels>pay_channel,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 赠券，格式为 券ID_券金额， 金额单位为分
	GiftCoupon string `json:"gift_coupon,omitempty" xml:"gift_coupon,omitempty"`
	// 买家id对应的open_uid，加密过
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 扩展属性map
	TradeAttributes string `json:"trade_attributes,omitempty" xml:"trade_attributes,omitempty"`
	// 收银班次号
	DutyCode string `json:"duty_code,omitempty" xml:"duty_code,omitempty"`
	// 打包时间
	PackageTime string `json:"package_time,omitempty" xml:"package_time,omitempty"`
	// 会员积分
	MemberPoint string `json:"member_point,omitempty" xml:"member_point,omitempty"`
	// 订单状态 PAID：已支付  PACKAGED:已打包 SUCCESS：已收货，订单完成
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 供应商code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 收银员名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 收银员id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 商场code
	SourceMerchantCode string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
	// 下单终端: APP,POS
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 淘宝订单号
	TbBizOrderId int64 `json:"tb_biz_order_id,omitempty" xml:"tb_biz_order_id,omitempty"`
	// 会员优惠金额
	MemberDiscountAmt int64 `json:"member_discount_amt,omitempty" xml:"member_discount_amt,omitempty"`
	// 邮费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 商家总原价
	OriginalAmt int64 `json:"original_amt,omitempty" xml:"original_amt,omitempty"`
	// 总优惠金额
	DiscountAmt int64 `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
	// 盒马订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolOrderSyncDto = sync.Pool{
	New: func() any {
		return new(OrderSyncDto)
	},
}

// GetOrderSyncDto() 从对象池中获取OrderSyncDto
func GetOrderSyncDto() *OrderSyncDto {
	return poolOrderSyncDto.Get().(*OrderSyncDto)
}

// ReleaseOrderSyncDto 释放OrderSyncDto
func ReleaseOrderSyncDto(v *OrderSyncDto) {
	v.SubOrders = v.SubOrders[:0]
	v.PayChannels = v.PayChannels[:0]
	v.ShopId = ""
	v.OutOrderId = ""
	v.GiftCoupon = ""
	v.OpenUid = ""
	v.TradeAttributes = ""
	v.DutyCode = ""
	v.PackageTime = ""
	v.MemberPoint = ""
	v.OrderStatus = ""
	v.MerchantCode = ""
	v.OperatorName = ""
	v.OperatorId = ""
	v.MemberCardNum = ""
	v.StoreId = ""
	v.PayTime = ""
	v.SourceMerchantCode = ""
	v.OrderClient = ""
	v.TbBizOrderId = 0
	v.MemberDiscountAmt = 0
	v.PostFee = 0
	v.OriginalAmt = 0
	v.DiscountAmt = 0
	v.BizOrderId = 0
	v.OrderFrom = 0
	poolOrderSyncDto.Put(v)
}
