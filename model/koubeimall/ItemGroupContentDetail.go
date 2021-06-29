package koubeimall

// ItemGroupContentDetail 
type ItemGroupContentDetail struct {
    // 金额
    ContentAmount   string `json:"content_amount,omitempty" xml:"content_amount,omitempty"`
    // 单位，份，杯
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 规格
    Spec   string `json:"spec,omitempty" xml:"spec,omitempty"`
    // 份数
    ContentCount   int64 `json:"content_count,omitempty" xml:"content_count,omitempty"`
    // 名称，eg：红酒
    ContentName   string `json:"content_name,omitempty" xml:"content_name,omitempty"`
}
