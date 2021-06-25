package promotion

// ShopInfoVo 
type ShopInfoVo struct {

    // 商家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 店铺链接
    ShopUrl   string `json:"shop_url,omitempty"`

    // 店铺icon链接
    ShopIconUrl   string `json:"shop_icon_url,omitempty"`

    // 店铺名称
    ShopName   string `json:"shop_name,omitempty"`

    // 店铺id
    ShopId   int64 `json:"shop_id,omitempty"`

}
