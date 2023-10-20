package miniappopen

import (
	"sync"
)

// MiniAppInstantiateTemplateAppSimpleRequest 结构体
type MiniAppInstantiateTemplateAppSimpleRequest struct {
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 小部件模版id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 小部件模版版本
	TemplateVersion string `json:"template_version,omitempty" xml:"template_version,omitempty"`
}

var poolMiniAppInstantiateTemplateAppSimpleRequest = sync.Pool{
	New: func() any {
		return new(MiniAppInstantiateTemplateAppSimpleRequest)
	},
}

// GetMiniAppInstantiateTemplateAppSimpleRequest() 从对象池中获取MiniAppInstantiateTemplateAppSimpleRequest
func GetMiniAppInstantiateTemplateAppSimpleRequest() *MiniAppInstantiateTemplateAppSimpleRequest {
	return poolMiniAppInstantiateTemplateAppSimpleRequest.Get().(*MiniAppInstantiateTemplateAppSimpleRequest)
}

// ReleaseMiniAppInstantiateTemplateAppSimpleRequest 释放MiniAppInstantiateTemplateAppSimpleRequest
func ReleaseMiniAppInstantiateTemplateAppSimpleRequest(v *MiniAppInstantiateTemplateAppSimpleRequest) {
	v.Description = ""
	v.TemplateId = ""
	v.TemplateVersion = ""
	poolMiniAppInstantiateTemplateAppSimpleRequest.Put(v)
}
