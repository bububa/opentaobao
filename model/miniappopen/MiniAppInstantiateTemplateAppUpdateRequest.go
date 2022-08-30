package miniappopen

// MiniAppInstantiateTemplateAppUpdateRequest 结构体
type MiniAppInstantiateTemplateAppUpdateRequest struct {
	// 小部件实例id
	EntityId string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// 小部件模版id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 小部件模版版本
	TemplateVersion string `json:"template_version,omitempty" xml:"template_version,omitempty"`
}
