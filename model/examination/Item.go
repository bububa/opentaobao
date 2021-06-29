package examination

// Item 
type Item struct {
    // 体检标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 体检项描述
    Detail   string `json:"detail,omitempty" xml:"detail,omitempty"`
    // 子标题
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    // 体检项编码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // 体检组编码
    ItemGroupCode   string `json:"item_group_code,omitempty" xml:"item_group_code,omitempty"`
    // 体检组显示权重
    ItemGroupWeight   string `json:"item_group_weight,omitempty" xml:"item_group_weight,omitempty"`
}
