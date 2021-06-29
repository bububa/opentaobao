package product

// CacheChangeInfo 
type CacheChangeInfo struct {
    // 外部酒店code
    HotelCode   string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
    // 变更时间段
    TimeSpan   *TimeSpan `json:"time_span,omitempty" xml:"time_span,omitempty"`
    // 具体变更的商品
    ProductInfo   *ProductInfo `json:"product_info,omitempty" xml:"product_info,omitempty"`
}
