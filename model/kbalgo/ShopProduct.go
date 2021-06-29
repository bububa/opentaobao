package kbalgo

// ShopProduct 
type ShopProduct struct {

    // Product_info
    
    ProductInfos   []ProductInfo `json:"product_infos,omitempty" xml:"product_infos,omitempty"`
    

    // 优惠券信息
    
    CouponInfos   []CouponInfo `json:"coupon_infos,omitempty" xml:"coupon_infos,omitempty"`
    

}
