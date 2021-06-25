package traderate

// TmallRateTagDetail 
type TmallRateTagDetail struct {

    // 标签名称
    TagName   string `json:"tag_name,omitempty"`

    // 反应该标签的评价数量
    Count   int64 `json:"count,omitempty"`

    // 标签的极性：1正极 -1负极
    Posi   bool `json:"posi,omitempty"`

}
