package maitix

// OpenCombinePriceDto 
type OpenCombinePriceDto struct {
    // 子票品ID
    PriceId   int64 `json:"price_id,omitempty" xml:"price_id,omitempty"`
    // 子票品名称
    PriceName   string `json:"price_name,omitempty" xml:"price_name,omitempty"`
    // 子票品价格，单位分
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 套票的子票数量
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
}
