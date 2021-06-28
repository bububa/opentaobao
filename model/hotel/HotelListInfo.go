package hotel

// HotelListInfo 
/* model for simplify = false
type HotelListInfo struct {

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 城市code
    
    CityCode   int64 `json:"city_code,omitempty"`
    

    // 评论条数
    
    CommentCount   int64 `json:"comment_count,omitempty"`
    

    // 评论分数
    
    CommentScore   string `json:"comment_score,omitempty"`
    

    // 评论分数的描述文案
    
    CommentScoreDesc   string `json:"comment_score_desc,omitempty"`
    

    // 评论分数
    
    CommentSource   int64 `json:"comment_source,omitempty"`
    

    // 距离文案
    
    DistanceDesc   string `json:"distance_desc,omitempty"`
    

    // 英文名
    
    EnglishName   string `json:"english_name,omitempty"`
    

    // 酒店fax
    
    Fax   string `json:"fax,omitempty"`
    

    // 是否有信用住商品
    
    HasAliCreditItem   bool `json:"has_ali_credit_item,omitempty"`
    

    // 是否有全景图
    
    HasFullScenePicture   bool `json:"has_full_scene_picture,omitempty"`
    

    // 头图地址
    
    HeaderPicUrl   string `json:"header_pic_url,omitempty"`
    

    // 酒店标签
    
    Labels  struct {
        HotelLabelVo  []HotelLabelVo `json:"hotel_label_vo,omitempty"`
    } `json:"labels,omitempty"`
    

    // 酒店经度
    
    Latitude   string `json:"latitude,omitempty"`
    

    // 酒店纬度
    
    Longitude   string `json:"longitude,omitempty"`
    

    // 酒店名字
    
    Name   string `json:"name,omitempty"`
    

    // 酒店报价
    
    Price   int64 `json:"price,omitempty"`
    

    // 标准酒店id
    
    Shid   int64 `json:"shid,omitempty"`
    

    // 是否售完
    
    SoldOut   bool `json:"sold_out,omitempty"`
    

    // 酒店星级类型
    
    StarType   string `json:"star_type,omitempty"`
    

    // 酒店电话
    
    Tel   string `json:"tel,omitempty"`
    

    // 品牌Id
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 品牌名字
    
    BrandName   string `json:"brand_name,omitempty"`
    

    // 品牌英文名字
    
    BrandNameEnglish   string `json:"brand_name_english,omitempty"`
    

    // inRightMapHotelTitle
    
    InRightMapHotelTitle   string `json:"in_right_map_hotel_title,omitempty"`
    

    // 活动
    
    PromotionDescArrs  struct {
        String  []string `json:"string,omitempty"`
    } `json:"promotion_desc_arrs,omitempty"`
    

    // searchPOI
    
    SearchPOI  *struct {
        SearchPoi  *SearchPoi `json:"search_poi,omitempty"`
    } `json:"search_p_o_i,omitempty"`
    

    // service
    
    Services  struct {
        Option  []Option `json:"option,omitempty"`
    } `json:"services,omitempty"`
    

    // laterPayIcon
    
    LaterPayIcon   string `json:"later_pay_icon,omitempty"`
    

    // inventoryDesc
    
    InventoryDesc   string `json:"inventory_desc,omitempty"`
    

    // 酒店特色id
    
    FeatureHotelType   string `json:"feature_hotel_type,omitempty"`
    

    // 酒店特色名称
    
    FeatureHotelTypeName   string `json:"feature_hotel_type_name,omitempty"`
    

    // 商圈
    
    Areas  struct {
        Area  []Area `json:"area,omitempty"`
    } `json:"areas,omitempty"`
    

}
*/

// HotelListInfo 
type HotelListInfo struct {

    // 地址
    Address   string `json:"address,omitempty"`

    // 城市code
    CityCode   int64 `json:"city_code,omitempty"`

    // 评论条数
    CommentCount   int64 `json:"comment_count,omitempty"`

    // 评论分数
    CommentScore   string `json:"comment_score,omitempty"`

    // 评论分数的描述文案
    CommentScoreDesc   string `json:"comment_score_desc,omitempty"`

    // 评论分数
    CommentSource   int64 `json:"comment_source,omitempty"`

    // 距离文案
    DistanceDesc   string `json:"distance_desc,omitempty"`

    // 英文名
    EnglishName   string `json:"english_name,omitempty"`

    // 酒店fax
    Fax   string `json:"fax,omitempty"`

    // 是否有信用住商品
    HasAliCreditItem   bool `json:"has_ali_credit_item,omitempty"`

    // 是否有全景图
    HasFullScenePicture   bool `json:"has_full_scene_picture,omitempty"`

    // 头图地址
    HeaderPicUrl   string `json:"header_pic_url,omitempty"`

    // 酒店标签
    Labels   []HotelLabelVo `json:"labels,omitempty"`

    // 酒店经度
    Latitude   string `json:"latitude,omitempty"`

    // 酒店纬度
    Longitude   string `json:"longitude,omitempty"`

    // 酒店名字
    Name   string `json:"name,omitempty"`

    // 酒店报价
    Price   int64 `json:"price,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // 是否售完
    SoldOut   bool `json:"sold_out,omitempty"`

    // 酒店星级类型
    StarType   string `json:"star_type,omitempty"`

    // 酒店电话
    Tel   string `json:"tel,omitempty"`

    // 品牌Id
    BrandId   string `json:"brand_id,omitempty"`

    // 品牌名字
    BrandName   string `json:"brand_name,omitempty"`

    // 品牌英文名字
    BrandNameEnglish   string `json:"brand_name_english,omitempty"`

    // inRightMapHotelTitle
    InRightMapHotelTitle   string `json:"in_right_map_hotel_title,omitempty"`

    // 活动
    PromotionDescArrs   []string `json:"promotion_desc_arrs,omitempty"`

    // searchPOI
    SearchPOI   *SearchPoi `json:"search_p_o_i,omitempty"`

    // service
    Services   []Option `json:"services,omitempty"`

    // laterPayIcon
    LaterPayIcon   string `json:"later_pay_icon,omitempty"`

    // inventoryDesc
    InventoryDesc   string `json:"inventory_desc,omitempty"`

    // 酒店特色id
    FeatureHotelType   string `json:"feature_hotel_type,omitempty"`

    // 酒店特色名称
    FeatureHotelTypeName   string `json:"feature_hotel_type_name,omitempty"`

    // 商圈
    Areas   []Area `json:"areas,omitempty"`

}
