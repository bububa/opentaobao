package feedflow

// OptionDTO 
type OptionDTO struct {
    // 标签值
    OptionValue   string `json:"option_value,omitempty" xml:"option_value,omitempty"`
    // 选项名称
    OptionName   string `json:"option_name,omitempty" xml:"option_name,omitempty"`
    // 选项描述
    OptionDesc   string `json:"option_desc,omitempty" xml:"option_desc,omitempty"`
}
