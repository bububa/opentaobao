package hotel

// HotelTopStaticDetailsVo 
type HotelTopStaticDetailsVo struct {

    // 地址
    Address   string `json:"address,omitempty"`

    // 行政区商圈
    Area   string `json:"area,omitempty"`

    // 开业或装修描述
    Decorate   string `json:"decorate,omitempty"`

    // 英文名字
    EnglishName   string `json:"english_name,omitempty"`

    // 小图标展示
    Icons   []String `json:"icons,omitempty"`

    // 消息信息
    InfoMsg   string `json:"info_msg,omitempty"`

    // 酒店名字
    Name   string `json:"name,omitempty"`

    // 全景图
    Panoramas   []Panorama `json:"panoramas,omitempty"`

    // 图片信息
    PicUrls   *PicStringArrayDo `json:"pic_urls,omitempty"`

    // 评论数
    RateNumber   int64 `json:"rate_number,omitempty"`

    // 综合评分
    RateScore   string `json:"rate_score,omitempty"`

    // 酒店评分文案说明（超出预期，非常好，很好）
    RateScoreDesc   string `json:"rate_score_desc,omitempty"`

    // 地标纬度
    RefLatitude   string `json:"ref_latitude,omitempty"`

    // 地标经度
    RefLongitude   string `json:"ref_longitude,omitempty"`

    // 地标名
    RefPoiName   string `json:"ref_poi_name,omitempty"`

    // 酒店服务
    Services   []Number `json:"services,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // 星级
    Star   string `json:"star,omitempty"`

    // 电话
    Tel   string `json:"tel,omitempty"`

    // commentSource
    CommentSource   int64 `json:"comment_source,omitempty"`

    // 动态uv
    DynamicUvDesc   string `json:"dynamic_uv_desc,omitempty"`

    // 是否港澳台
    IsGangAoTai   int64 `json:"is_gang_ao_tai,omitempty"`

    // 是否国内
    IsInternational   int64 `json:"is_international,omitempty"`

    // 经度
    Latitude   string `json:"latitude,omitempty"`

    // 纬度
    Longitude   string `json:"longitude,omitempty"`

    // rateScoreText
    RateScoreText   string `json:"rate_score_text,omitempty"`

    // 得分描述
    ScoreText   string `json:"score_text,omitempty"`

    // 特征
    Feature   string `json:"feature,omitempty"`

    // 装修时间
    DecorateTime   string `json:"decorate_time,omitempty"`

    // 描述
    Description   string `json:"description,omitempty"`

    // 设施icon
    FacilitieIcons   []FacilityIcons `json:"facilitie_icons,omitempty"`

    // 设施
    Facilities   []String `json:"facilities,omitempty"`

    // 税
    Fax   string `json:"fax,omitempty"`

    // 图片
    Galleries   []Gallery `json:"galleries,omitempty"`

    // 开业时间
    OpeningTime   string `json:"opening_time,omitempty"`

    // 房屋设施
    RoomFacilities   []String `json:"room_facilities,omitempty"`

    // 设施信息 for pc
    Services4Pcs   []String `json:"services4_pcs,omitempty"`

    // hotelPolicyList
    HotelPolicyList   []String `json:"hotel_policy_list,omitempty"`

}
