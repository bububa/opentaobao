package icbudropshipping

// FreightTemplateRequest 结构体
type FreightTemplateRequest struct {
	// destination country  ISO 3166-2
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// destination zip code
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
}
