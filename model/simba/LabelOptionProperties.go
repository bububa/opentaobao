package simba

import (
	"sync"
)

// LabelOptionProperties 结构体
type LabelOptionProperties struct {
	// dmp选项分组id
	DmpOptionGroupId string `json:"dmp_option_group_id,omitempty" xml:"dmp_option_group_id,omitempty"`
	// 物料id
	CateName string `json:"cate_name,omitempty" xml:"cate_name,omitempty"`
	// 物料id
	CateId string `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
}

var poolLabelOptionProperties = sync.Pool{
	New: func() any {
		return new(LabelOptionProperties)
	},
}

// GetLabelOptionProperties() 从对象池中获取LabelOptionProperties
func GetLabelOptionProperties() *LabelOptionProperties {
	return poolLabelOptionProperties.Get().(*LabelOptionProperties)
}

// ReleaseLabelOptionProperties 释放LabelOptionProperties
func ReleaseLabelOptionProperties(v *LabelOptionProperties) {
	v.DmpOptionGroupId = ""
	v.CateName = ""
	v.CateId = ""
	poolLabelOptionProperties.Put(v)
}
