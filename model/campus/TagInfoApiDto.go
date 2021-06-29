package campus

// TagInfoApiDTO 
type TagInfoApiDTO struct {
    // 标签名称
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    // 是否系统标签
    SystemTag   bool `json:"system_tag,omitempty" xml:"system_tag,omitempty"`
    // 标签描述
    TagDesc   string `json:"tag_desc,omitempty" xml:"tag_desc,omitempty"`
    // 标签类型
    TagTypeName   string `json:"tag_type_name,omitempty" xml:"tag_type_name,omitempty"`
}
