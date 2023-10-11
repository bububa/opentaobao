package simba

// PriceVo 结构体
type PriceVo struct {
	// 出价
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 建议出价
	FitBidPrice string `json:"fit_bid_price,omitempty" xml:"fit_bid_price,omitempty"`
	// 溢价
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 建议溢价
	FitDiscount string `json:"fit_discount,omitempty" xml:"fit_discount,omitempty"`
}
