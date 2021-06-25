package product

// ItemSearchResult 
type ItemSearchResult struct {

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品标题
    ItemTitle   string `json:"item_title,omitempty"`

    // 商品链接
    ItemUrl   string `json:"item_url,omitempty"`

    // 商品主图
    ItemMainPic   string `json:"item_main_pic,omitempty"`

    // 商品原价最低值
    ItemOriginPriceMin   string `json:"item_origin_price_min,omitempty"`

    // 商品原价最高值
    ItemOriginPriceMax   string `json:"item_origin_price_max,omitempty"`

    // 商品折扣价最低值
    ItemPriceDiscountMin   string `json:"item_price_discount_min,omitempty"`

    // 商品折扣价最高值
    ItemPriceDiscountMax   string `json:"item_price_discount_max,omitempty"`

    // 商品折扣率
    ItemDiscountRate   string `json:"item_discount_rate,omitempty"`

    // 商品发布时间
    PubTime   string `json:"pub_time,omitempty"`

    // 商品评价分数
    CommentScore   string `json:"comment_score,omitempty"`

    // 店铺URL
    ShopUrl   string `json:"shop_url,omitempty"`

    // 佣金比例
    CommissionRate   string `json:"commission_rate,omitempty"`

    // 商品图片地址
    ItemPics   string `json:"item_pics,omitempty"`

    // 商品视频地址
    ItemVideos   string `json:"item_videos,omitempty"`

    // 卖家服务等级
    SellerLayer   string `json:"seller_layer,omitempty"`

    // 30天销量语义化信息
    Sales30DaySemantic   string `json:"sales30_day_semantic,omitempty"`

    // 30天评论数语义化信息
    Comment30DaySemantic   string `json:"comment30_day_semantic,omitempty"`

    // 收藏数语义化信息
    FavoriteCntSemantic   string `json:"favorite_cnt_semantic,omitempty"`

}
