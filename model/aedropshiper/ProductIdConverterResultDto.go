package aedropshiper

// ProductIdConverterResultDto 结构体
type ProductIdConverterResultDto struct {
	// sub productId
	SubProductId string `json:"sub_product_id,omitempty" xml:"sub_product_id,omitempty"`
	// main productId
	MainProductId int64 `json:"main_product_id,omitempty" xml:"main_product_id,omitempty"`
}
