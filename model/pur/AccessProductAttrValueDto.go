package pur

// AccessProductAttrValueDto 结构体
type AccessProductAttrValueDto struct {
	// 属性Value值英文
	AttrEnValueNameList []string `json:"attr_en_value_name_list,omitempty" xml:"attr_en_value_name_list>string,omitempty"`
	// 属性Value值中文
	AttrValueNameList []string `json:"attr_value_name_list,omitempty" xml:"attr_value_name_list>string,omitempty"`
	// 属性Key英文值
	AttrEnName string `json:"attr_en_name,omitempty" xml:"attr_en_name,omitempty"`
	// 属性Key中文值
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// 属性描述
	AttrDesc string `json:"attr_desc,omitempty" xml:"attr_desc,omitempty"`
	// 是否关键属性y/n
	IsKey string `json:"is_key,omitempty" xml:"is_key,omitempty"`
}
