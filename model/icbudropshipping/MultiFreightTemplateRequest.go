package icbudropshipping

// MultiFreightTemplateRequest 
type MultiFreightTemplateRequest struct {
    // Get from alibaba.dropshipping.product.get
    ECompanyId   string `json:"e_company_id,omitempty" xml:"e_company_id,omitempty"`
    // Shipping address
    Address   *AddressInfoDTO `json:"address,omitempty" xml:"address,omitempty"`
    // Destination Country
    DestinationCountry   string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
    // Product List
    LogisticsProductList   []LogisticsProduct `json:"logistics_product_list,omitempty" xml:"logistics_product_list>logistics_product,omitempty"`
}
