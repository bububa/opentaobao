package hotel

// RoomTypeInfo 
type RoomTypeInfo struct {

    // 面积
    Acreage   string `json:"acreage,omitempty"`

    // 床型
    BedType   string `json:"bed_type,omitempty"`

    // 英文名
    EnglishName   string `json:"english_name,omitempty"`

    // 楼层
    Floor   string `json:"floor,omitempty"`

    // 展平商品
    Items   []ItemInfo `json:"items,omitempty"`

    // 只存标签id
    Labels   []Number `json:"labels,omitempty"`

    // 房型名称
    Name   string `json:"name,omitempty"`

    // 可住人数
    Occupancy   string `json:"occupancy,omitempty"`

    // 全景房
    Panoramas   []Panorama `json:"panoramas,omitempty"`

    // 图片信息
    PicUrls   *PicStringArrayDo `json:"pic_urls,omitempty"`

    // 最低价
    Price   int64 `json:"price,omitempty"`

    // 房型设施
    Services   string `json:"services,omitempty"`

    // 标准化房型ID
    Srtid   int64 `json:"srtid,omitempty"`

    // 窗型信息
    WindowType   string `json:"window_type,omitempty"`

    // 是否还有更多
    IsMore   int64 `json:"is_more,omitempty"`

    // 是否满房
    IsSellOut   int64 `json:"is_sell_out,omitempty"`

    // backCash
    BackCash   *model.File `json:"back_cash,omitempty"`

    // drid
    Drid   string `json:"drid,omitempty"`

    // firstStay
    FirstStay   bool `json:"first_stay,omitempty"`

    // 是否热销
    HotSale   bool `json:"hot_sale,omitempty"`

    // 立减状态：0:无立减标，1：显示立减标
    ImmediatelySubtract   int64 `json:"immediately_subtract,omitempty"`

    // laterPay
    LaterPay   bool `json:"later_pay,omitempty"`

    // memberPrice
    MemberPrice   bool `json:"member_price,omitempty"`

    // networkService
    NetworkService   string `json:"network_service,omitempty"`

    // 大促展示的文案内容
    PromotionDescArrs   []String `json:"promotion_desc_arrs,omitempty"`

    // 立减
    SubtractPrice   int64 `json:"subtract_price,omitempty"`

    // b2bVip
    B2bVip   bool `json:"b2b_vip,omitempty"`

    // bedTypes
    BedTypes   []String `json:"bed_types,omitempty"`

    // double12Desc
    Double12Desc   string `json:"double12_desc,omitempty"`

    // dualEleven
    DualEleven   bool `json:"dual_eleven,omitempty"`

    // showNewPeopleCash
    ShowNewPeopleCash   bool `json:"show_new_people_cash,omitempty"`

    // priceDesc
    PriceDesc   string `json:"price_desc,omitempty"`

}
