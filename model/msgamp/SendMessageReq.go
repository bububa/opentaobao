package msgamp

import (
	"sync"
)

// SendMessageReq 结构体
type SendMessageReq struct {
	// 接受者ID，如果是群，则是群ID。如果是用户，则是uid。
	TargetId string `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 选用的模板ID
	TemplateInstanceId string `json:"template_instance_id,omitempty" xml:"template_instance_id,omitempty"`
	// 模板渲染参数
	TemplateData string `json:"template_data,omitempty" xml:"template_data,omitempty"`
	// 跳转的目标链接参数，如：page=xxxx
	ActionUrlParams string `json:"action_url_params,omitempty" xml:"action_url_params,omitempty"`
}

var poolSendMessageReq = sync.Pool{
	New: func() any {
		return new(SendMessageReq)
	},
}

// GetSendMessageReq() 从对象池中获取SendMessageReq
func GetSendMessageReq() *SendMessageReq {
	return poolSendMessageReq.Get().(*SendMessageReq)
}

// ReleaseSendMessageReq 释放SendMessageReq
func ReleaseSendMessageReq(v *SendMessageReq) {
	v.TargetId = ""
	v.TemplateInstanceId = ""
	v.TemplateData = ""
	v.ActionUrlParams = ""
	poolSendMessageReq.Put(v)
}
