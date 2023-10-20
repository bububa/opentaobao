package jst

import (
	"sync"
)

// TopDeleteSmsTemplateRequest 结构体
type TopDeleteSmsTemplateRequest struct {
	// 待删除的模板code
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
}

var poolTopDeleteSmsTemplateRequest = sync.Pool{
	New: func() any {
		return new(TopDeleteSmsTemplateRequest)
	},
}

// GetTopDeleteSmsTemplateRequest() 从对象池中获取TopDeleteSmsTemplateRequest
func GetTopDeleteSmsTemplateRequest() *TopDeleteSmsTemplateRequest {
	return poolTopDeleteSmsTemplateRequest.Get().(*TopDeleteSmsTemplateRequest)
}

// ReleaseTopDeleteSmsTemplateRequest 释放TopDeleteSmsTemplateRequest
func ReleaseTopDeleteSmsTemplateRequest(v *TopDeleteSmsTemplateRequest) {
	v.TemplateCode = ""
	poolTopDeleteSmsTemplateRequest.Put(v)
}
