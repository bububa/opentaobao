package wlbimports

import (
	"sync"
)

// AppointmentCreateRequest 结构体
type AppointmentCreateRequest struct {
	// 预约大包信息列表（（非自寄模式必填）
	HandoverContentSynopsisList []OrderHandoverContentSynopsisDto `json:"handover_content_synopsis_list,omitempty" xml:"handover_content_synopsis_list>order_handover_content_synopsis_dto,omitempty"`
	// 小包信息（自寄模式必填）
	LgOrderList []string `json:"lg_order_list,omitempty" xml:"lg_order_list>string,omitempty"`
	// 时区
	ZoneOffSet string `json:"zone_off_set,omitempty" xml:"zone_off_set,omitempty"`
	// 接收仓资源名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 揽收方式：TRUCK（卡车）OFFLINE_EXPRESS(自寄快递) OFFLINE_TRUCK(自派卡车)
	PickupType string `json:"pickup_type,omitempty" xml:"pickup_type,omitempty"`
	// 接收仓资源编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 发货人信息
	SenderInfo *ContactInfoRequest `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 预约大包信息列表（自送快递和自送卡车非必填）
	ReceiverInfo *ContactInfoRequest `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 自寄快递请求参数（自寄快递模式必填）
	SelfSendExpressInfoRequest *SelfSendExpressInfoRequest `json:"self_send_express_info_request,omitempty" xml:"self_send_express_info_request,omitempty"`
	// 揽收信息（（非自寄模式必填）
	PickupInfo *PickupInfo `json:"pickup_info,omitempty" xml:"pickup_info,omitempty"`
	// 自寄卡车参数（自寄卡车模式必填）
	SelfSendTruckInfoRequest *SelfMailTruckInfoRequest `json:"self_send_truck_info_request,omitempty" xml:"self_send_truck_info_request,omitempty"`
}

var poolAppointmentCreateRequest = sync.Pool{
	New: func() any {
		return new(AppointmentCreateRequest)
	},
}

// GetAppointmentCreateRequest() 从对象池中获取AppointmentCreateRequest
func GetAppointmentCreateRequest() *AppointmentCreateRequest {
	return poolAppointmentCreateRequest.Get().(*AppointmentCreateRequest)
}

// ReleaseAppointmentCreateRequest 释放AppointmentCreateRequest
func ReleaseAppointmentCreateRequest(v *AppointmentCreateRequest) {
	v.HandoverContentSynopsisList = v.HandoverContentSynopsisList[:0]
	v.LgOrderList = v.LgOrderList[:0]
	v.ZoneOffSet = ""
	v.StoreName = ""
	v.PickupType = ""
	v.StoreCode = ""
	v.SellerId = 0
	v.SenderInfo = nil
	v.ReceiverInfo = nil
	v.SelfSendExpressInfoRequest = nil
	v.PickupInfo = nil
	v.SelfSendTruckInfoRequest = nil
	poolAppointmentCreateRequest.Put(v)
}
