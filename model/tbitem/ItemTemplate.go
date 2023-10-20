package tbitem

import (
	"sync"
)

// ItemTemplate 结构体
type ItemTemplate struct {
	// 宝贝详情模板的名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 宝贝模板的id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 用于区分宝贝模板属于内店和外店
	ShopType int64 `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
}

var poolItemTemplate = sync.Pool{
	New: func() any {
		return new(ItemTemplate)
	},
}

// GetItemTemplate() 从对象池中获取ItemTemplate
func GetItemTemplate() *ItemTemplate {
	return poolItemTemplate.Get().(*ItemTemplate)
}

// ReleaseItemTemplate 释放ItemTemplate
func ReleaseItemTemplate(v *ItemTemplate) {
	v.TemplateName = ""
	v.TemplateId = 0
	v.ShopType = 0
	poolItemTemplate.Put(v)
}
