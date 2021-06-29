package cainiaohandover

// OpenPickupInfoParam 
type OpenPickupInfoParam struct {
    // 卖家后台地址id,用来获取卖家详细地址信息，传入值为Long型；
    SellerAddressId   int64 `json:"seller_address_id,omitempty" xml:"seller_address_id,omitempty"`
}
