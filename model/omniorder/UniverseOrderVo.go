package omniorder

import (
	"sync"
)

// UniverseOrderVo 结构体
type UniverseOrderVo struct {
	// 订单支付时间：yyyy-MM-dd HH:mm:ss
	OrderPayTime string `json:"order_pay_time,omitempty" xml:"order_pay_time,omitempty"`
	// 描述订单的唯一值，如订单id、加密后的订单id等，用于订单幂等去重
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 订单创建时间：yyyy-MM-dd HH:mm:ss
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 导购工号
	GuideWorkId string `json:"guide_work_id,omitempty" xml:"guide_work_id,omitempty"`
	// 订单结束时间：yyyy-MM-dd HH:mm:ss
	OrderEndTime string `json:"order_end_time,omitempty" xml:"order_end_time,omitempty"`
	// 订单发货时间：yyyy-MM-dd HH:mm:ss
	OrderGoodsTime string `json:"order_goods_time,omitempty" xml:"order_goods_time,omitempty"`
	// 导购手机号
	GuidePhone string `json:"guide_phone,omitempty" xml:"guide_phone,omitempty"`
	// 消费者nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 商品id
	AuctionId string `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 子订单id
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 消费者手机号
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 导购id
	EmployeeId int64 `json:"employee_id,omitempty" xml:"employee_id,omitempty"`
	// 分佣金额，单位分
	CommissionMoney int64 `json:"commission_money,omitempty" xml:"commission_money,omitempty"`
	// 订单类型： 1-线上 2-线下门店 3-云店 4-其他
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 订单渠道：1-天猫旗舰店  2-门店pos  3-云店  4-京东 5-微信商城 6-官网 7-美团 999-其他
	OrderSource int64 `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 订单状态 0：支付（默认） 1：发货 2：确认收货
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 消费者id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 当前订单的分佣比例。如1%，则填1
	CommissionRatio int64 `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
	// 订单发生时的销售额，单位分
	GuideSales int64 `json:"guide_sales,omitempty" xml:"guide_sales,omitempty"`
	// 订单金额，单位分
	OrderMoney int64 `json:"order_money,omitempty" xml:"order_money,omitempty"`
}

var poolUniverseOrderVo = sync.Pool{
	New: func() any {
		return new(UniverseOrderVo)
	},
}

// GetUniverseOrderVo() 从对象池中获取UniverseOrderVo
func GetUniverseOrderVo() *UniverseOrderVo {
	return poolUniverseOrderVo.Get().(*UniverseOrderVo)
}

// ReleaseUniverseOrderVo 释放UniverseOrderVo
func ReleaseUniverseOrderVo(v *UniverseOrderVo) {
	v.OrderPayTime = ""
	v.OuterOrderId = ""
	v.OrderCreateTime = ""
	v.GuideWorkId = ""
	v.OrderEndTime = ""
	v.OrderGoodsTime = ""
	v.GuidePhone = ""
	v.BuyerNick = ""
	v.AuctionId = ""
	v.BizOrderId = ""
	v.OrderId = ""
	v.BuyerPhone = ""
	v.EmployeeId = 0
	v.CommissionMoney = 0
	v.OrderType = 0
	v.OrderSource = 0
	v.OrderStatus = 0
	v.BuyerId = 0
	v.CommissionRatio = 0
	v.GuideSales = 0
	v.OrderMoney = 0
	poolUniverseOrderVo.Put(v)
}
