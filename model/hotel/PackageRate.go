package hotel

// PackageRate 
type PackageRate struct {

    // 面积
    Acreage   string `json:"acreage,omitempty"`

    // 床型描述
    BedDesc   string `json:"bed_desc,omitempty"`

    // 床型
    BedType   string `json:"bed_type,omitempty"`

    // 早餐
    Breakfast   string `json:"breakfast,omitempty"`

    // 英文名字
    EnglishName   string `json:"english_name,omitempty"`

    // 楼层
    Floor   string `json:"floor,omitempty"`

    // 酒店id
    Hid   int64 `json:"hid,omitempty"`

    // 高亮内容
    HighlightTitle   *HighlightContent `json:"highlight_title,omitempty"`

    // 打包count
    HotelPackageCount   int64 `json:"hotel_package_count,omitempty"`

    // 酒店描述
    HotelPackageDesc   string `json:"hotel_package_desc,omitempty"`

    // 酒店图片
    HotelPackagePic   string `json:"hotel_package_pic,omitempty"`

    // 打包title
    HotelPackageTitle   string `json:"hotel_package_title,omitempty"`

    // 酒店商品ID
    Id   int64 `json:"id,omitempty"`

    // IC商品数字ID
    Iid   int64 `json:"iid,omitempty"`

    // 库存描述
    InventoryDesc   string `json:"inventory_desc,omitempty"`

    // 是否担保商品
    IsGuarantee   int64 `json:"is_guarantee,omitempty"`

    // 是否是打包
    IsHotelPackage   int64 `json:"is_hotel_package,omitempty"`

    // 是否是会员 0否1是
    IsMember   int64 `json:"is_member,omitempty"`

    // 是否卖完
    IsSellOut   int64 `json:"is_sell_out,omitempty"`

    // 只存标签id
    Labels   []Number `json:"labels,omitempty"`

    // 房型名称
    Name   string `json:"name,omitempty"`

    // 连住人数
    Nop   int64 `json:"nop,omitempty"`

    // 可住人数
    Occupancy   string `json:"occupancy,omitempty"`

    // 预定链接
    OrderConfirmUnits   *OrderConfirmUnits `json:"order_confirm_units,omitempty"`

    // 预订成功率
    OrderSucessRate   int64 `json:"order_sucess_rate,omitempty"`

    // 打包信息
    PackageInfos   []HotelPackageVo `json:"package_infos,omitempty"`

    // 全景图
    Panoramas   []Panorama `json:"panoramas,omitempty"`

    // 支付类型。编码取值：1:全额支付;5:前台面付;
    PaymentType   int64 `json:"payment_type,omitempty"`

    // 房型图片名称
    PicUrls   *PicStringArrayDo `json:"pic_urls,omitempty"`

    // rate上的价格描述文案
    PriceDesc   string `json:"price_desc,omitempty"`

    // rateId
    RateId   int64 `json:"rate_id,omitempty"`

    // 推荐文案
    RecommendDesc   string `json:"recommend_desc,omitempty"`

    // 退订信息
    RefundInfo   string `json:"refund_info,omitempty"`

    // 退订规则
    RefundRules   string `json:"refund_rules,omitempty"`

    // 房型内容类型
    RoomContents   []HighlightContent `json:"room_contents,omitempty"`

    // 商品定价规则标题
    RpTitle   string `json:"rp_title,omitempty"`

    // 商品定价规则ID，类似SKU
    Rpid   int64 `json:"rpid,omitempty"`

    // 卖家房型id
    RtId   int64 `json:"rt_id,omitempty"`

    // 卖家ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 卖家综合评分
    SellerScore   string `json:"seller_score,omitempty"`

    // 房型设施
    Services   string `json:"services,omitempty"`

    // 展示价格（6.1.0版后为减后价）
    ShowPrice   int64 `json:"show_price,omitempty"`

    // 标准化房型ID
    Srtid   int64 `json:"srtid,omitempty"`

    // 窗型
    WindowType   string `json:"window_type,omitempty"`

    // breakfastDesc
    BreakfastDesc   string `json:"breakfast_desc,omitempty"`

    // breakfasts
    Breakfasts   []String `json:"breakfasts,omitempty"`

    // costPrice
    CostPrice   string `json:"cost_price,omitempty"`

    // firstStayLimits
    FirstStayLimits   []Number `json:"first_stay_limits,omitempty"`

    // futures
    Futures   []Number `json:"futures,omitempty"`

    // hidden
    Hidden   string `json:"hidden,omitempty"`

    // orderShipTime
    OrderShipTime   int64 `json:"order_ship_time,omitempty"`

    // rtEnglishName
    RtEnglishName   string `json:"rt_english_name,omitempty"`

    // rtName
    RtName   string `json:"rt_name,omitempty"`

    // sellerScoreThanAvg
    SellerScoreThanAvg   int64 `json:"seller_score_than_avg,omitempty"`

    // shipTimeThanAvg
    ShipTimeThanAvg   int64 `json:"ship_time_than_avg,omitempty"`

    // subTitle
    SubTitle   string `json:"sub_title,omitempty"`

    // successRateThanAvg
    SuccessRateThanAvg   int64 `json:"success_rate_than_avg,omitempty"`

    // taxAndFee
    TaxAndFee   string `json:"tax_and_fee,omitempty"`

}
