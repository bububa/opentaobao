package jst

import (
	"sync"
)

// SendMessageByOaidRequest 结构体
type SendMessageByOaidRequest struct {
	// 拓展Name
	ExtendName string `json:"extend_name,omitempty" xml:"extend_name,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 拓展Code
	ExtendCode string `json:"extend_code,omitempty" xml:"extend_code,omitempty"`
	// 短信签名
	SmsFreeSignName string `json:"sms_free_sign_name,omitempty" xml:"sms_free_sign_name,omitempty"`
	// 短信模板
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 短信占位符，替换短信模板里的占位符
	Params string `json:"params,omitempty" xml:"params,omitempty"`
	// 拓展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// OAID
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 是否需要返回附加信息
	AdditionalInfo bool `json:"additional_info,omitempty" xml:"additional_info,omitempty"`
}

var poolSendMessageByOaidRequest = sync.Pool{
	New: func() any {
		return new(SendMessageByOaidRequest)
	},
}

// GetSendMessageByOaidRequest() 从对象池中获取SendMessageByOaidRequest
func GetSendMessageByOaidRequest() *SendMessageByOaidRequest {
	return poolSendMessageByOaidRequest.Get().(*SendMessageByOaidRequest)
}

// ReleaseSendMessageByOaidRequest 释放SendMessageByOaidRequest
func ReleaseSendMessageByOaidRequest(v *SendMessageByOaidRequest) {
	v.ExtendName = ""
	v.OrderId = ""
	v.ExtendCode = ""
	v.SmsFreeSignName = ""
	v.TemplateCode = ""
	v.Params = ""
	v.Extend = ""
	v.Oaid = ""
	v.AdditionalInfo = false
	poolSendMessageByOaidRequest.Put(v)
}
