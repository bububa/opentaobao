package wdk

// Couponconfiglist 
type Couponconfiglist struct {
    // 提货券面额，单位：元，用于展示
    ViewTotalAmount   string `json:"view_total_amount,omitempty" xml:"view_total_amount,omitempty"`
    // 提货券面额，单位：元，用于展示
    ViewAmount   string `json:"view_amount,omitempty" xml:"view_amount,omitempty"`
    // 提货券总金额（单位：分）
    UseTotalAmount   int64 `json:"use_total_amount,omitempty" xml:"use_total_amount,omitempty"`
    // 面额相应数量
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 提货券面额（单位：分）
    UseAmount   int64 `json:"use_amount,omitempty" xml:"use_amount,omitempty"`
}
