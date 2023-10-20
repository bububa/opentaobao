package tbtrade

import (
	"sync"
)

// DeliveryPlan 结构体
type DeliveryPlan struct {
	// 外部物流单号
	OutLogisticsId string `json:"out_logistics_id,omitempty" xml:"out_logistics_id,omitempty"`
	// 备货期开始时间
	PrepareTimeBegin string `json:"prepare_time_begin,omitempty" xml:"prepare_time_begin,omitempty"`
	// 计划发货时间
	ShipTimeBegin string `json:"ship_time_begin,omitempty" xml:"ship_time_begin,omitempty"`
	// 实际发货时间
	ActualShipTime string `json:"actual_ship_time,omitempty" xml:"actual_ship_time,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人手机号
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收货人电话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收货详细地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// aid
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 隐私号
	PrivacyNum string `json:"privacy_num,omitempty" xml:"privacy_num,omitempty"`
	// 隐私号过期时间
	PrivacyExpireTime string `json:"privacy_expire_time,omitempty" xml:"privacy_expire_time,omitempty"`
	// 收货地址
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
	// 收货地址
	ReceiverDistrict string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
	// 收货地址
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 收货地址
	ReceiverState string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
	// 收货地址
	ReceiverCountry string `json:"receiver_country,omitempty" xml:"receiver_country,omitempty"`
	// 计划id
	PlanId int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// 交易订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 发货期数
	CurrPhase int64 `json:"curr_phase,omitempty" xml:"curr_phase,omitempty"`
	// 发货商品数量
	GoodsNum int64 `json:"goods_num,omitempty" xml:"goods_num,omitempty"`
	// 计划状态
	PlanStatus int64 `json:"plan_status,omitempty" xml:"plan_status,omitempty"`
	// 计划退款状态
	PlanRefundStatus int64 `json:"plan_refund_status,omitempty" xml:"plan_refund_status,omitempty"`
}

var poolDeliveryPlan = sync.Pool{
	New: func() any {
		return new(DeliveryPlan)
	},
}

// GetDeliveryPlan() 从对象池中获取DeliveryPlan
func GetDeliveryPlan() *DeliveryPlan {
	return poolDeliveryPlan.Get().(*DeliveryPlan)
}

// ReleaseDeliveryPlan 释放DeliveryPlan
func ReleaseDeliveryPlan(v *DeliveryPlan) {
	v.OutLogisticsId = ""
	v.PrepareTimeBegin = ""
	v.ShipTimeBegin = ""
	v.ActualShipTime = ""
	v.ReceiverName = ""
	v.ReceiverMobile = ""
	v.ReceiverPhone = ""
	v.ReceiverAddress = ""
	v.Oaid = ""
	v.PrivacyNum = ""
	v.PrivacyExpireTime = ""
	v.ReceiverTown = ""
	v.ReceiverDistrict = ""
	v.ReceiverCity = ""
	v.ReceiverState = ""
	v.ReceiverCountry = ""
	v.PlanId = 0
	v.OrderId = 0
	v.CurrPhase = 0
	v.GoodsNum = 0
	v.PlanStatus = 0
	v.PlanRefundStatus = 0
	poolDeliveryPlan.Put(v)
}
