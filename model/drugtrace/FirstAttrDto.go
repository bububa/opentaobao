package drugtrace

// FirstAttrDto 结构体
type FirstAttrDto struct {
	// 一级药物属性名称
	FirstAttributeName string `json:"first_attribute_name,omitempty" xml:"first_attribute_name,omitempty"`
	// 一级药物属性编号
	FirstAttributeNo string `json:"first_attribute_no,omitempty" xml:"first_attribute_no,omitempty"`
	// 二级药物属性信息
	SecondaryAttrDtoList []SecondaryAttrDto `json:"secondary_attr_dto_list,omitempty" xml:"secondary_attr_dto_list>secondary_attr_dto,omitempty"`
}
