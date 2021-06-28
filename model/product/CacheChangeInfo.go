package product

// CacheChangeInfo 
/* model for simplify = false
type CacheChangeInfo struct {

    // 外部酒店code
    
    HotelCode   string `json:"hotel_code,omitempty"`
    

    // 变更时间段
    
    TimeSpan  *struct {
        TimeSpan  *TimeSpan `json:"time_span,omitempty"`
    } `json:"time_span,omitempty"`
    

    // 具体变更的商品
    
    ProductInfo  *struct {
        ProductInfo  *ProductInfo `json:"product_info,omitempty"`
    } `json:"product_info,omitempty"`
    

}
*/

// CacheChangeInfo 
type CacheChangeInfo struct {

    // 外部酒店code
    HotelCode   string `json:"hotel_code,omitempty"`

    // 变更时间段
    TimeSpan   *TimeSpan `json:"time_span,omitempty"`

    // 具体变更的商品
    ProductInfo   *ProductInfo `json:"product_info,omitempty"`

}
