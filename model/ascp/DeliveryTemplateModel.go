package ascp

import (
	"sync"
)

// DeliveryTemplateModel 结构体
type DeliveryTemplateModel struct {
	// 运费模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 运费模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolDeliveryTemplateModel = sync.Pool{
	New: func() any {
		return new(DeliveryTemplateModel)
	},
}

// GetDeliveryTemplateModel() 从对象池中获取DeliveryTemplateModel
func GetDeliveryTemplateModel() *DeliveryTemplateModel {
	return poolDeliveryTemplateModel.Get().(*DeliveryTemplateModel)
}

// ReleaseDeliveryTemplateModel 释放DeliveryTemplateModel
func ReleaseDeliveryTemplateModel(v *DeliveryTemplateModel) {
	v.Name = ""
	v.TemplateId = 0
	poolDeliveryTemplateModel.Put(v)
}
