package fenxiao

// RegionalPriceDto 
type RegionalPriceDto struct {

    // 市
    City   string `json:"city,omitempty"`

    // 金额（分）
    Price   int64 `json:"price,omitempty"`

    // 省
    Province   string `json:"province,omitempty"`

    // 区县，特殊可选
    District   string `json:"district,omitempty"`

    // 街道，特殊可选
    Street   string `json:"street,omitempty"`

}
