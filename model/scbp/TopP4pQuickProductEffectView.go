package scbp

// TopP4pQuickProductEffectView 
type TopP4pQuickProductEffectView struct {

    // 产品名称
    ProductName   string `json:"product_name,omitempty"`

    // 产品id
    ProductId   int64 `json:"product_id,omitempty"`

    // 平均点击花费
    Cpc   string `json:"cpc,omitempty"`

    // 点击率
    Ctr   string `json:"ctr,omitempty"`

    // 消耗金额，单位元
    Cost   string `json:"cost,omitempty"`

    // 点击量
    ClickCnt   string `json:"click_cnt,omitempty"`

    // 曝光量
    ImpressionCnt   string `json:"impression_cnt,omitempty"`

}
