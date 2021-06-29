package aedropshiper

// AeopStoreInfo 
type AeopStoreInfo struct {
    // 卖家服务，1-5星
    CommunicationRating   string `json:"communication_rating,omitempty" xml:"communication_rating,omitempty"`
    // 商品描述，1-5星
    ItemAsDescripedRating   string `json:"item_as_descriped_rating,omitempty" xml:"item_as_descriped_rating,omitempty"`
    // 物流，1-5星
    ShippingSpeedRating   string `json:"shipping_speed_rating,omitempty" xml:"shipping_speed_rating,omitempty"`
    // 店铺ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 店铺名
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 店铺地址
    StoreUrl   string `json:"store_url,omitempty" xml:"store_url,omitempty"`
}
