package damai

// ProjectDto 
type ProjectDto struct {

    // 演出城市
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 演出时间
    
    ShowTime   string `json:"show_time,omitempty" xml:"show_time,omitempty"`
    

    // 项目主类
    
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    

    // 场次开始时间
    
    PerformStartTime   string `json:"perform_start_time,omitempty" xml:"perform_start_time,omitempty"`
    

    // 项目图片地址
    
    VerticalPic   string `json:"vertical_pic,omitempty" xml:"vertical_pic,omitempty"`
    

    // 项目名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 艺人数据
    
    Actors   string `json:"actors,omitempty" xml:"actors,omitempty"`
    

    // 场馆城市
    
    VenueCity   string `json:"venue_city,omitempty" xml:"venue_city,omitempty"`
    

    // 结束售卖时间
    
    SellEndTime   int64 `json:"sell_end_time,omitempty" xml:"sell_end_time,omitempty"`
    

    // 场馆名称
    
    VenueName   string `json:"venue_name,omitempty" xml:"venue_name,omitempty"`
    

    // 上架时间
    
    UpTime   string `json:"up_time,omitempty" xml:"up_time,omitempty"`
    

    // 关联艺人
    
    ArtistName   string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
    

    // 项目子类
    
    SubCategoryName   string `json:"sub_category_name,omitempty" xml:"sub_category_name,omitempty"`
    

    // 巡演信息
    
    Tours   string `json:"tours,omitempty" xml:"tours,omitempty"`
    

    // 关联品牌
    
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    

    // 商品状态 3:售票中
    
    SiteStatus   string `json:"site_status,omitempty" xml:"site_status,omitempty"`
    

    // 推荐语
    
    SubHead   string `json:"sub_head,omitempty" xml:"sub_head,omitempty"`
    

    // 扩展信息
    
    ExtraInfoMap   *Extrainfomap `json:"extra_info_map,omitempty" xml:"extra_info_map,omitempty"`
    

    // 是否电子票 1:是，0:非
    
    IsETicket   string `json:"is_e_ticket,omitempty" xml:"is_e_ticket,omitempty"`
    

    // 是否选座 1:是,0:非
    
    IsSelectSeat   string `json:"is_select_seat,omitempty" xml:"is_select_seat,omitempty"`
    

    // 优惠价格
    
    PromotionPrice   string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
    

    // 短标题
    
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    

    // 价格字符串
    
    PriceStr   string `json:"price_str,omitempty" xml:"price_str,omitempty"`
    

    // 演出场馆id
    
    VenueId   int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
    

    // 纬度
    
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    

    // 经度
    
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    

}
