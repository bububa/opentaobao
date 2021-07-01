package tmallservice

// StoreOfferPriceDto 
type StoreOfferPriceDto struct {
    // price
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 服务code
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    // 服务skuCode
    ServiceSku   string `json:"service_sku,omitempty" xml:"service_sku,omitempty"`
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
