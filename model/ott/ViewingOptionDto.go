package ott

// ViewingOptionDto 结构体
type ViewingOptionDto struct {
	// currency
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// license
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// quality
	Quality string `json:"quality,omitempty" xml:"quality,omitempty"`
}
