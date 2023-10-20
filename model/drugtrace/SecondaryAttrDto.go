package drugtrace

import (
	"sync"
)

// SecondaryAttrDto 结构体
type SecondaryAttrDto struct {
	// 二级药物属性名称
	SecondaryAttributeName string `json:"secondary_attribute_name,omitempty" xml:"secondary_attribute_name,omitempty"`
	// 二级药物属性编号
	SecondaryAttributeNo string `json:"secondary_attribute_no,omitempty" xml:"secondary_attribute_no,omitempty"`
	// 二级药物属性描述
	SecondaryAttrDesc string `json:"secondary_attr_desc,omitempty" xml:"secondary_attr_desc,omitempty"`
	// 药品id
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
}

var poolSecondaryAttrDto = sync.Pool{
	New: func() any {
		return new(SecondaryAttrDto)
	},
}

// GetSecondaryAttrDto() 从对象池中获取SecondaryAttrDto
func GetSecondaryAttrDto() *SecondaryAttrDto {
	return poolSecondaryAttrDto.Get().(*SecondaryAttrDto)
}

// ReleaseSecondaryAttrDto 释放SecondaryAttrDto
func ReleaseSecondaryAttrDto(v *SecondaryAttrDto) {
	v.SecondaryAttributeName = ""
	v.SecondaryAttributeNo = ""
	v.SecondaryAttrDesc = ""
	v.DrugEntBaseInfoId = ""
	poolSecondaryAttrDto.Put(v)
}
