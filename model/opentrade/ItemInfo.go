package opentrade

// ItemInfo 
type ItemInfo struct {

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品SKU
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 商品数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 订单页显示的商品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 订单页显示的商品主图地址，只能为图片空间中的图片。只需传入图片地址的path，无需域名
    
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    

}
