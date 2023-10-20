package alsc

import (
	"sync"
)

// TemplateNotifyRequest 结构体
type TemplateNotifyRequest struct {
	// 用户OpenId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 小程序appid
	Appid string `json:"appid,omitempty" xml:"appid,omitempty"`
	// 消息主体码
	NotifyId string `json:"notify_id,omitempty" xml:"notify_id,omitempty"`
	// 业务标签
	BizTag string `json:"biz_tag,omitempty" xml:"biz_tag,omitempty"`
	// 模板变量数据
	TemplateArgs string `json:"template_args,omitempty" xml:"template_args,omitempty"`
}

var poolTemplateNotifyRequest = sync.Pool{
	New: func() any {
		return new(TemplateNotifyRequest)
	},
}

// GetTemplateNotifyRequest() 从对象池中获取TemplateNotifyRequest
func GetTemplateNotifyRequest() *TemplateNotifyRequest {
	return poolTemplateNotifyRequest.Get().(*TemplateNotifyRequest)
}

// ReleaseTemplateNotifyRequest 释放TemplateNotifyRequest
func ReleaseTemplateNotifyRequest(v *TemplateNotifyRequest) {
	v.OpenId = ""
	v.Appid = ""
	v.NotifyId = ""
	v.BizTag = ""
	v.TemplateArgs = ""
	poolTemplateNotifyRequest.Put(v)
}
