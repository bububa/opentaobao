package traveltrade

import (
	"sync"
)

// OrderTemplateInfo 结构体
type OrderTemplateInfo struct {
	// 填充字段列表
	ModelList []OrderTipFormInfo `json:"model_list,omitempty" xml:"model_list>order_tip_form_info,omitempty"`
	// 模版对应的类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 模版对应ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 模版对应版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOrderTemplateInfo = sync.Pool{
	New: func() any {
		return new(OrderTemplateInfo)
	},
}

// GetOrderTemplateInfo() 从对象池中获取OrderTemplateInfo
func GetOrderTemplateInfo() *OrderTemplateInfo {
	return poolOrderTemplateInfo.Get().(*OrderTemplateInfo)
}

// ReleaseOrderTemplateInfo 释放OrderTemplateInfo
func ReleaseOrderTemplateInfo(v *OrderTemplateInfo) {
	v.ModelList = v.ModelList[:0]
	v.CategoryId = 0
	v.TemplateId = 0
	v.Version = 0
	poolOrderTemplateInfo.Put(v)
}
