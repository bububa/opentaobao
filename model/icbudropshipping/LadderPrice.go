package icbudropshipping

// LadderPrice 结构体
type LadderPrice struct {
	// price
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// If it is -1, it means the maximum
	MaxQuantity int64 `json:"max_quantity,omitempty" xml:"max_quantity,omitempty"`
	// min quantity
	MinQuantity int64 `json:"min_quantity,omitempty" xml:"min_quantity,omitempty"`
}
