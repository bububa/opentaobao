package drugtrace

import (
	"sync"
)

// FirstAttrDto 结构体
type FirstAttrDto struct {
	// 二级药物属性信息
	SecondaryAttrDtoList []SecondaryAttrDto `json:"secondary_attr_dto_list,omitempty" xml:"secondary_attr_dto_list>secondary_attr_dto,omitempty"`
	// 一级药物属性名称
	FirstAttributeName string `json:"first_attribute_name,omitempty" xml:"first_attribute_name,omitempty"`
	// 一级药物属性编号
	FirstAttributeNo string `json:"first_attribute_no,omitempty" xml:"first_attribute_no,omitempty"`
}

var poolFirstAttrDto = sync.Pool{
	New: func() any {
		return new(FirstAttrDto)
	},
}

// GetFirstAttrDto() 从对象池中获取FirstAttrDto
func GetFirstAttrDto() *FirstAttrDto {
	return poolFirstAttrDto.Get().(*FirstAttrDto)
}

// ReleaseFirstAttrDto 释放FirstAttrDto
func ReleaseFirstAttrDto(v *FirstAttrDto) {
	v.SecondaryAttrDtoList = v.SecondaryAttrDtoList[:0]
	v.FirstAttributeName = ""
	v.FirstAttributeNo = ""
	poolFirstAttrDto.Put(v)
}
