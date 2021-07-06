package icbudropshipping

// MultiFreightTemplateRequest 结构体
type MultiFreightTemplateRequest struct {
	// Product List
	LogisticsProductList []LogisticsProduct `json:"logistics_product_list,omitempty" xml:"logistics_product_list>logistics_product,omitempty"`
	// Get from alibaba.dropshipping.product.get
	ECompanyId string `json:"e_company_id,omitempty" xml:"e_company_id,omitempty"`
	// Destination Country
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// Shipping address
	Address *AddressInfoDto `json:"address,omitempty" xml:"address,omitempty"`
}
