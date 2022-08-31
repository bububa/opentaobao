package logistic

// ProductDto 结构体
type ProductDto struct {
	// Total weight of a SKU in its original packaging. Default unit: kilograms
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// Actual dimension of a sku in its original packaging. Default unit: centimeters
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// Price of product
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// AE sku_id to identify a unit of sku
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// Quantity of a sku in the order. It&#39;s used to calculate the total number of products in a parcel
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
