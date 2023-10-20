package logistic

import (
	"sync"
)

// NotifyCutOffDeliveryProcessStatusRequest 结构体
type NotifyCutOffDeliveryProcessStatusRequest struct {
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 业务单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 拦截状态，1拦截成功，2拦截失败，3拦截取消
	InterceptStatus string `json:"intercept_status,omitempty" xml:"intercept_status,omitempty"`
	// 快递公司
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 拦截失败、拦截取消原因
	InterceptStatusMessage string `json:"intercept_status_message,omitempty" xml:"intercept_status_message,omitempty"`
}

var poolNotifyCutOffDeliveryProcessStatusRequest = sync.Pool{
	New: func() any {
		return new(NotifyCutOffDeliveryProcessStatusRequest)
	},
}

// GetNotifyCutOffDeliveryProcessStatusRequest() 从对象池中获取NotifyCutOffDeliveryProcessStatusRequest
func GetNotifyCutOffDeliveryProcessStatusRequest() *NotifyCutOffDeliveryProcessStatusRequest {
	return poolNotifyCutOffDeliveryProcessStatusRequest.Get().(*NotifyCutOffDeliveryProcessStatusRequest)
}

// ReleaseNotifyCutOffDeliveryProcessStatusRequest 释放NotifyCutOffDeliveryProcessStatusRequest
func ReleaseNotifyCutOffDeliveryProcessStatusRequest(v *NotifyCutOffDeliveryProcessStatusRequest) {
	v.MailNo = ""
	v.OuterOrderId = ""
	v.InterceptStatus = ""
	v.TmsCpCode = ""
	v.InterceptStatusMessage = ""
	poolNotifyCutOffDeliveryProcessStatusRequest.Put(v)
}
