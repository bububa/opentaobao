package logistic

import (
	"sync"
)

// Tms2MscUpdateOrderRequest 结构体
type Tms2MscUpdateOrderRequest struct {
	// 电联信息（组）
	PhoneCallInfos []TmsPhoneCallInfoDto `json:"phone_call_infos,omitempty" xml:"phone_call_infos>tms_phone_call_info_dto,omitempty"`
	// 业务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 修改后的上门日期（YYYY-MM-DD）
	ServiceDate string `json:"service_date,omitempty" xml:"service_date,omitempty"`
	// 请求类型
	RequestType string `json:"request_type,omitempty" xml:"request_type,omitempty"`
	// 唯一标识单号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 修改上门时间原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 修改后的上门时间段（HH:MM-HH:MM）
	ServiceTimeRange string `json:"service_time_range,omitempty" xml:"service_time_range,omitempty"`
	// 配资源编码
	DeliverCode string `json:"deliver_code,omitempty" xml:"deliver_code,omitempty"`
}

var poolTms2MscUpdateOrderRequest = sync.Pool{
	New: func() any {
		return new(Tms2MscUpdateOrderRequest)
	},
}

// GetTms2MscUpdateOrderRequest() 从对象池中获取Tms2MscUpdateOrderRequest
func GetTms2MscUpdateOrderRequest() *Tms2MscUpdateOrderRequest {
	return poolTms2MscUpdateOrderRequest.Get().(*Tms2MscUpdateOrderRequest)
}

// ReleaseTms2MscUpdateOrderRequest 释放Tms2MscUpdateOrderRequest
func ReleaseTms2MscUpdateOrderRequest(v *Tms2MscUpdateOrderRequest) {
	v.PhoneCallInfos = v.PhoneCallInfos[:0]
	v.ServiceType = ""
	v.SupplierId = ""
	v.ServiceDate = ""
	v.RequestType = ""
	v.BizCode = ""
	v.Reason = ""
	v.ServiceTimeRange = ""
	v.DeliverCode = ""
	poolTms2MscUpdateOrderRequest.Put(v)
}
