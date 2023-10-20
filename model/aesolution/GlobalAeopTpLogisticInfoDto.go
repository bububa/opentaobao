package aesolution

import (
	"sync"
)

// GlobalAeopTpLogisticInfoDto 结构体
type GlobalAeopTpLogisticInfoDto struct {
	// logistics tracking number
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// un-receive reason,such as Country does not match
	RecvStatusDesc string `json:"recv_status_desc,omitempty" xml:"recv_status_desc,omitempty"`
	// logistics service show name
	LogisticsServiceName string `json:"logistics_service_name,omitempty" xml:"logistics_service_name,omitempty"`
	// send time
	GmtSend string `json:"gmt_send,omitempty" xml:"gmt_send,omitempty"`
	// received time
	GmtReceived string `json:"gmt_received,omitempty" xml:"gmt_received,omitempty"`
	// logistics service name key
	LogisticsTypeCode string `json:"logistics_type_code,omitempty" xml:"logistics_type_code,omitempty"`
	// receive status。(default:initial value; received:; not_received; suspected_received)
	ReceiveStatus string `json:"receive_status,omitempty" xml:"receive_status,omitempty"`
	// ship order id
	ShipOrderId int64 `json:"ship_order_id,omitempty" xml:"ship_order_id,omitempty"`
	// to get logistics tracking information
	HaveTrackingInfo bool `json:"have_tracking_info,omitempty" xml:"have_tracking_info,omitempty"`
}

var poolGlobalAeopTpLogisticInfoDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpLogisticInfoDto)
	},
}

// GetGlobalAeopTpLogisticInfoDto() 从对象池中获取GlobalAeopTpLogisticInfoDto
func GetGlobalAeopTpLogisticInfoDto() *GlobalAeopTpLogisticInfoDto {
	return poolGlobalAeopTpLogisticInfoDto.Get().(*GlobalAeopTpLogisticInfoDto)
}

// ReleaseGlobalAeopTpLogisticInfoDto 释放GlobalAeopTpLogisticInfoDto
func ReleaseGlobalAeopTpLogisticInfoDto(v *GlobalAeopTpLogisticInfoDto) {
	v.LogisticsNo = ""
	v.RecvStatusDesc = ""
	v.LogisticsServiceName = ""
	v.GmtSend = ""
	v.GmtReceived = ""
	v.LogisticsTypeCode = ""
	v.ReceiveStatus = ""
	v.ShipOrderId = 0
	v.HaveTrackingInfo = false
	poolGlobalAeopTpLogisticInfoDto.Put(v)
}
