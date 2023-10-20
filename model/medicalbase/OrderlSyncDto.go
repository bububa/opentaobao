package medicalbase

import (
	"sync"
)

// OrderlSyncDto 结构体
type OrderlSyncDto struct {
	// 登录用户支付宝ID
	AlipayId string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// 医院唯一标识
	HosOrgNo string `json:"hos_org_no,omitempty" xml:"hos_org_no,omitempty"`
	// 院区编码
	HosDistinctCode string `json:"hos_distinct_code,omitempty" xml:"hos_distinct_code,omitempty"`
	// 预约ID
	ReservationId string `json:"reservation_id,omitempty" xml:"reservation_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 失败原因
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 取号密码
	OrderPass string `json:"order_pass,omitempty" xml:"order_pass,omitempty"`
}

var poolOrderlSyncDto = sync.Pool{
	New: func() any {
		return new(OrderlSyncDto)
	},
}

// GetOrderlSyncDto() 从对象池中获取OrderlSyncDto
func GetOrderlSyncDto() *OrderlSyncDto {
	return poolOrderlSyncDto.Get().(*OrderlSyncDto)
}

// ReleaseOrderlSyncDto 释放OrderlSyncDto
func ReleaseOrderlSyncDto(v *OrderlSyncDto) {
	v.AlipayId = ""
	v.HosOrgNo = ""
	v.HosDistinctCode = ""
	v.ReservationId = ""
	v.OrderStatus = ""
	v.Remark = ""
	v.OrderPass = ""
	poolOrderlSyncDto.Put(v)
}
