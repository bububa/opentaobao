package scbp

// Effect7d 
type Effect7d struct {
    // 点击率
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    // 平均点击花费
    Cpc   string `json:"cpc,omitempty" xml:"cpc,omitempty"`
    // 消耗金额，单位元
    Cost   string `json:"cost,omitempty" xml:"cost,omitempty"`
    // 点击量
    Click   string `json:"click,omitempty" xml:"click,omitempty"`
    // 曝光量
    Impr   string `json:"impr,omitempty" xml:"impr,omitempty"`
}
