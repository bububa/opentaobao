package idle

// IdleItemApiDo 
type IdleItemApiDo struct {

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 商品类目
    
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 渠道类目ID
    
    ChannelCatId   string `json:"channel_cat_id,omitempty" xml:"channel_cat_id,omitempty"`
    

    // 商品图片
    
    ImgUrls   []string `json:"img_urls,omitempty" xml:"img_urls>string,omitempty"`
    

    // 商品价格
    
    ReservePrice   string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
    

}
