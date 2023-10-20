package alihouse

import (
	"sync"
)

// BaseLabelDto 结构体
type BaseLabelDto struct {
	// 父级业务
	ParentBusiness string `json:"parent_business,omitempty" xml:"parent_business,omitempty"`
	// 业务
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 分类
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 标签ID
	LabelId string `json:"label_id,omitempty" xml:"label_id,omitempty"`
	// 标签名称
	LabelName string `json:"label_name,omitempty" xml:"label_name,omitempty"`
	// 类型ID
	TypeId string `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 业务ID
	BusinessId string `json:"business_id,omitempty" xml:"business_id,omitempty"`
	// 父级业务ID
	ParentBusinessId string `json:"parent_business_id,omitempty" xml:"parent_business_id,omitempty"`
	// 业务类型0-新房 1-二手房
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolBaseLabelDto = sync.Pool{
	New: func() any {
		return new(BaseLabelDto)
	},
}

// GetBaseLabelDto() 从对象池中获取BaseLabelDto
func GetBaseLabelDto() *BaseLabelDto {
	return poolBaseLabelDto.Get().(*BaseLabelDto)
}

// ReleaseBaseLabelDto 释放BaseLabelDto
func ReleaseBaseLabelDto(v *BaseLabelDto) {
	v.ParentBusiness = ""
	v.Business = ""
	v.Type = ""
	v.LabelId = ""
	v.LabelName = ""
	v.TypeId = ""
	v.BusinessId = ""
	v.ParentBusinessId = ""
	v.BizType = 0
	poolBaseLabelDto.Put(v)
}
