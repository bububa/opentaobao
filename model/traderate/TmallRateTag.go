package traderate

// TmallRateTag 
type TmallRateTag struct {
    // 表示标签的名称
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
    // 表示标签的极性，正极true，负极false
    Posi   bool `json:"posi,omitempty" xml:"posi,omitempty"`
}
