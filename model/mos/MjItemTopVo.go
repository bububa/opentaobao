package mos

// MjItemTopVo 
/* model for simplify = false
type MjItemTopVo struct {

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 是否独享365会员
    
    ItemTag365   bool `json:"item_tag365,omitempty"`
    

    // 活动标签
    
    MarketingActivityTag   string `json:"marketing_activity_tag,omitempty"`
    

    // 最小促销价
    
    MinPromotionPrice   int64 `json:"min_promotion_price,omitempty"`
    

    // 商品图片，数组中只有一个值
    
    Pics  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pics,omitempty"`
    

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // minPrice
    
    MinPrice   int64 `json:"min_price,omitempty"`
    

    // maxPrice
    
    MaxPrice   int64 `json:"max_price,omitempty"`
    

    // 活动限定人群，1：普通人群，2：365会员
    
    RequiredUserType   int64 `json:"required_user_type,omitempty"`
    

}
*/

// MjItemTopVo 
type MjItemTopVo struct {

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 是否独享365会员
    ItemTag365   bool `json:"item_tag365,omitempty"`

    // 活动标签
    MarketingActivityTag   string `json:"marketing_activity_tag,omitempty"`

    // 最小促销价
    MinPromotionPrice   int64 `json:"min_promotion_price,omitempty"`

    // 商品图片，数组中只有一个值
    Pics   []string `json:"pics,omitempty"`

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // minPrice
    MinPrice   int64 `json:"min_price,omitempty"`

    // maxPrice
    MaxPrice   int64 `json:"max_price,omitempty"`

    // 活动限定人群，1：普通人群，2：365会员
    RequiredUserType   int64 `json:"required_user_type,omitempty"`

}
