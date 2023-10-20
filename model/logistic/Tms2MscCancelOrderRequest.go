package logistic

import (
	"sync"
)

// Tms2MscCancelOrderRequest 结构体
type Tms2MscCancelOrderRequest struct {
	// 电联信息（组）
	PhoneCallInfos []TmsPhoneCallInfoDto `json:"phone_call_infos,omitempty" xml:"phone_call_infos>tms_phone_call_info_dto,omitempty"`
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 唯一标识单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 取消上门订单原因，枚举
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 配资源编码
	DeliverCode string `json:"deliver_code,omitempty" xml:"deliver_code,omitempty"`
}

var poolTms2MscCancelOrderRequest = sync.Pool{
	New: func() any {
		return new(Tms2MscCancelOrderRequest)
	},
}

// GetTms2MscCancelOrderRequest() 从对象池中获取Tms2MscCancelOrderRequest
func GetTms2MscCancelOrderRequest() *Tms2MscCancelOrderRequest {
	return poolTms2MscCancelOrderRequest.Get().(*Tms2MscCancelOrderRequest)
}

// ReleaseTms2MscCancelOrderRequest 释放Tms2MscCancelOrderRequest
func ReleaseTms2MscCancelOrderRequest(v *Tms2MscCancelOrderRequest) {
	v.PhoneCallInfos = v.PhoneCallInfos[:0]
	v.ServiceType = ""
	v.SupplierId = ""
	v.BizCode = ""
	v.Reason = ""
	v.DeliverCode = ""
	poolTms2MscCancelOrderRequest.Put(v)
}
