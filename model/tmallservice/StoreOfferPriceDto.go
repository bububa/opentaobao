package tmallservice

// StoreOfferPriceDto 
/* model for simplify = false
type StoreOfferPriceDto struct {

    // price
    
    Price   int64 `json:"price,omitempty"`
    

    // 服务code
    
    ServiceCode   string `json:"service_code,omitempty"`
    

    // 服务skuCode
    
    ServiceSku   string `json:"service_sku,omitempty"`
    

    // 门店ID
    
    StoreId   int64 `json:"store_id,omitempty"`
    

}
*/

// StoreOfferPriceDto 
type StoreOfferPriceDto struct {

    // price
    Price   int64 `json:"price,omitempty"`

    // 服务code
    ServiceCode   string `json:"service_code,omitempty"`

    // 服务skuCode
    ServiceSku   string `json:"service_sku,omitempty"`

    // 门店ID
    StoreId   int64 `json:"store_id,omitempty"`

}
