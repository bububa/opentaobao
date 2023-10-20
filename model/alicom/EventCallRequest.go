package alicom

import (
	"sync"
)

// EventCallRequest 结构体
type EventCallRequest struct {
	// 唯一的呼叫ID，最大可支持字符串长度256
	CallId string `json:"call_id,omitempty" xml:"call_id,omitempty"`
	// 被叫号码
	CalledNo string `json:"called_no,omitempty" xml:"called_no,omitempty"`
	// 主叫号码
	CallNo string `json:"call_no,omitempty" xml:"call_no,omitempty"`
	// 分机号
	ExtensionNo string `json:"extension_no,omitempty" xml:"extension_no,omitempty"`
	// 振铃事件：ALERTING  摘机事件：PICKUP  被叫呼出事件：CALLOUT
	EventType string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	// 绑定关系ID
	SubsId string `json:"subs_id,omitempty" xml:"subs_id,omitempty"`
	// 供应商KEY
	VendorKey string `json:"vendor_key,omitempty" xml:"vendor_key,omitempty"`
	// 中间号码
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 事件时间
	EventTime string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 被叫号显
	CalledDisplayNo string `json:"called_display_no,omitempty" xml:"called_display_no,omitempty"`
	// 呼叫开始时间
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 呼叫前转号码
	CallForwardingNo string `json:"call_forwarding_no,omitempty" xml:"call_forwarding_no,omitempty"`
}

var poolEventCallRequest = sync.Pool{
	New: func() any {
		return new(EventCallRequest)
	},
}

// GetEventCallRequest() 从对象池中获取EventCallRequest
func GetEventCallRequest() *EventCallRequest {
	return poolEventCallRequest.Get().(*EventCallRequest)
}

// ReleaseEventCallRequest 释放EventCallRequest
func ReleaseEventCallRequest(v *EventCallRequest) {
	v.CallId = ""
	v.CalledNo = ""
	v.CallNo = ""
	v.ExtensionNo = ""
	v.EventType = ""
	v.SubsId = ""
	v.VendorKey = ""
	v.SecretNo = ""
	v.EventTime = ""
	v.CalledDisplayNo = ""
	v.CallTime = ""
	v.CallForwardingNo = ""
	poolEventCallRequest.Put(v)
}
