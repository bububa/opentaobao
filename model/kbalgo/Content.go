package kbalgo

// Content 
type Content struct {

    // 到家信息
    
    HomeProduct   *HomeProduct `json:"home_product,omitempty" xml:"home_product,omitempty"`
    

    // Poi
    
    Poi   *Poi `json:"poi,omitempty" xml:"poi,omitempty"`
    

    // 到店信息
    
    ShopProduct   *ShopProduct `json:"shop_product,omitempty" xml:"shop_product,omitempty"`
    

}
