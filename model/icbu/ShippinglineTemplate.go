package icbu

import (
	"sync"
)

// ShippinglineTemplate 结构体
type ShippinglineTemplate struct {
	// 运费模板名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 运费模板id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolShippinglineTemplate = sync.Pool{
	New: func() any {
		return new(ShippinglineTemplate)
	},
}

// GetShippinglineTemplate() 从对象池中获取ShippinglineTemplate
func GetShippinglineTemplate() *ShippinglineTemplate {
	return poolShippinglineTemplate.Get().(*ShippinglineTemplate)
}

// ReleaseShippinglineTemplate 释放ShippinglineTemplate
func ReleaseShippinglineTemplate(v *ShippinglineTemplate) {
	v.Title = ""
	v.Id = 0
	poolShippinglineTemplate.Put(v)
}
