package idleisv

// IdleNewPubPropertyValueDo 
type IdleNewPubPropertyValueDo struct {
    // 属性下所有的值
    ValuesList   []IdleNewPubValueDo `json:"values_list,omitempty" xml:"values_list>idle_new_pub_value_do,omitempty"`
    // 属性的名称，显示
    PropertyName   string `json:"property_name,omitempty" xml:"property_name,omitempty"`
    // 是否多选
    IsMultiple   bool `json:"is_multiple,omitempty" xml:"is_multiple,omitempty"`
    // 输入提示文本
    InputWord   string `json:"input_word,omitempty" xml:"input_word,omitempty"`
    // 属性的key
    PropertyId   string `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
