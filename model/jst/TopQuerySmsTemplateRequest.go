package jst

import (
	"sync"
)

// TopQuerySmsTemplateRequest 结构体
type TopQuerySmsTemplateRequest struct {
	// 要查询的模板CODE
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
}

var poolTopQuerySmsTemplateRequest = sync.Pool{
	New: func() any {
		return new(TopQuerySmsTemplateRequest)
	},
}

// GetTopQuerySmsTemplateRequest() 从对象池中获取TopQuerySmsTemplateRequest
func GetTopQuerySmsTemplateRequest() *TopQuerySmsTemplateRequest {
	return poolTopQuerySmsTemplateRequest.Get().(*TopQuerySmsTemplateRequest)
}

// ReleaseTopQuerySmsTemplateRequest 释放TopQuerySmsTemplateRequest
func ReleaseTopQuerySmsTemplateRequest(v *TopQuerySmsTemplateRequest) {
	v.TemplateCode = ""
	poolTopQuerySmsTemplateRequest.Put(v)
}
