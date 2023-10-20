package aedropshiper

import (
	"sync"
)

// AeOrderInfoResultDto 结构体
type AeOrderInfoResultDto struct {
	// Order logistics information list
	LogisticsInfoList []AeOrderLogisticsInfo `json:"logistics_info_list,omitempty" xml:"logistics_info_list>ae_order_logistics_info,omitempty"`
	// Sub-order list
	ChildOrderList []AeChildOrderInfo `json:"child_order_list,omitempty" xml:"child_order_list>ae_child_order_info,omitempty"`
	// Order creation time
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// Logistics status
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// Order Status
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// Order amount
	OrderAmount *SimpleMoney `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// Store Information
	StoreInfo *AeStoreSimpleInfo `json:"store_info,omitempty" xml:"store_info,omitempty"`
}

var poolAeOrderInfoResultDto = sync.Pool{
	New: func() any {
		return new(AeOrderInfoResultDto)
	},
}

// GetAeOrderInfoResultDto() 从对象池中获取AeOrderInfoResultDto
func GetAeOrderInfoResultDto() *AeOrderInfoResultDto {
	return poolAeOrderInfoResultDto.Get().(*AeOrderInfoResultDto)
}

// ReleaseAeOrderInfoResultDto 释放AeOrderInfoResultDto
func ReleaseAeOrderInfoResultDto(v *AeOrderInfoResultDto) {
	v.LogisticsInfoList = v.LogisticsInfoList[:0]
	v.ChildOrderList = v.ChildOrderList[:0]
	v.GmtCreate = ""
	v.LogisticsStatus = ""
	v.OrderStatus = ""
	v.OrderAmount = nil
	v.StoreInfo = nil
	poolAeOrderInfoResultDto.Put(v)
}
