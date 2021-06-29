package drugtrace

// ProducePreAttributeDTO 
type ProducePreAttributeDTO struct {
    // 货品属性对象
    AttrInfoList   []Attrinfolist `json:"attr_info_list,omitempty" xml:"attr_info_list>attrinfolist,omitempty"`
    // 属性规则-英文
    DefaultProducePreAttributeEn   string `json:"default_produce_pre_attribute_en,omitempty" xml:"default_produce_pre_attribute_en,omitempty"`
    // 属性规则-中文
    DefaultProducePreAttribute   string `json:"default_produce_pre_attribute,omitempty" xml:"default_produce_pre_attribute,omitempty"`
}
