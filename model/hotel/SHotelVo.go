package hotel

// SHotelVo 
/* model for simplify = false
type SHotelVo struct {

    // 地址
    
    Address   string `json:"address,omitempty"`
    

    // 是否是民宿
    
    BnbHotel   bool `json:"bnb_hotel,omitempty"`
    

    // 品牌code
    
    Brand   string `json:"brand,omitempty"`
    

    // 最早入住时间
    
    CheckInTime   string `json:"check_in_time,omitempty"`
    

    // 最晚离店时间
    
    CheckOutTime   string `json:"check_out_time,omitempty"`
    

    // 当前酒店所在城市
    
    City   int64 `json:"city,omitempty"`
    

    // 装修时间
    
    DecorateTime   string `json:"decorate_time,omitempty"`
    

    // 酒店描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 与某个POI的距离
    
    Distance   int64 `json:"distance,omitempty"`
    

    // 区县
    
    District   int64 `json:"district,omitempty"`
    

    // 最后修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // h5详情页url
    
    H5DetailUrl   string `json:"h5_detail_url,omitempty"`
    

    // 默认PC详情页url
    
    HotelDetailUrl   string `json:"hotel_detail_url,omitempty"`
    

    // 酒店设施文案
    
    HotelFacilities   string `json:"hotel_facilities,omitempty"`
    

    // 维度
    
    Lat   string `json:"lat,omitempty"`
    

    // 档次
    
    Level  *struct {
        NameValuePair  *NameValuePair `json:"name_value_pair,omitempty"`
    } `json:"level,omitempty"`
    

    // 经度
    
    Lng   string `json:"lng,omitempty"`
    

    // 酒店中文名
    
    Name   string `json:"name,omitempty"`
    

    // 开业时间
    
    OpeningTime   string `json:"opening_time,omitempty"`
    

    // 重复信息
    
    PetInfo   string `json:"pet_info,omitempty"`
    

    // 图片信息
    
    PicUrls   string `json:"pic_urls,omitempty"`
    

    // 价格
    
    Price   int64 `json:"price,omitempty"`
    

    // 省份
    
    Province   int64 `json:"province,omitempty"`
    

    // 评论数
    
    RateNumber   int64 `json:"rate_number,omitempty"`
    

    // 评分
    
    RateScore   string `json:"rate_score,omitempty"`
    

    // 销量
    
    Sell   int64 `json:"sell,omitempty"`
    

    // 标准酒店id
    
    Shid   int64 `json:"shid,omitempty"`
    

    // 服务设施集合类
    
    ShotelPropertiesVo  *struct {
        ShotelPropertiesSetVo  *ShotelPropertiesSetVo `json:"shotel_properties_set_vo,omitempty"`
    } `json:"shotel_properties_vo,omitempty"`
    

    // 标准房型列表
    
    SroomTypes  struct {
        SRoomTypeVo  []SRoomTypeVo `json:"s_room_type_vo,omitempty"`
    } `json:"sroom_types,omitempty"`
    

    // 档次，0--无档次，2--二星及经济，3--舒适，4--高档，5--豪华
    
    Star   string `json:"star,omitempty"`
    

    // 上下架状态，0--下架，其他状态-下架
    
    Status   int64 `json:"status,omitempty"`
    

    // 电话
    
    Tel   string `json:"tel,omitempty"`
    

    // 酒店业务分类
    
    Type   string `json:"type,omitempty"`
    

    // 客房数
    
    Rooms   int64 `json:"rooms,omitempty"`
    

}
*/

// SHotelVo 
type SHotelVo struct {

    // 地址
    Address   string `json:"address,omitempty"`

    // 是否是民宿
    BnbHotel   bool `json:"bnb_hotel,omitempty"`

    // 品牌code
    Brand   string `json:"brand,omitempty"`

    // 最早入住时间
    CheckInTime   string `json:"check_in_time,omitempty"`

    // 最晚离店时间
    CheckOutTime   string `json:"check_out_time,omitempty"`

    // 当前酒店所在城市
    City   int64 `json:"city,omitempty"`

    // 装修时间
    DecorateTime   string `json:"decorate_time,omitempty"`

    // 酒店描述
    Desc   string `json:"desc,omitempty"`

    // 与某个POI的距离
    Distance   int64 `json:"distance,omitempty"`

    // 区县
    District   int64 `json:"district,omitempty"`

    // 最后修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // h5详情页url
    H5DetailUrl   string `json:"h5_detail_url,omitempty"`

    // 默认PC详情页url
    HotelDetailUrl   string `json:"hotel_detail_url,omitempty"`

    // 酒店设施文案
    HotelFacilities   string `json:"hotel_facilities,omitempty"`

    // 维度
    Lat   string `json:"lat,omitempty"`

    // 档次
    Level   *NameValuePair `json:"level,omitempty"`

    // 经度
    Lng   string `json:"lng,omitempty"`

    // 酒店中文名
    Name   string `json:"name,omitempty"`

    // 开业时间
    OpeningTime   string `json:"opening_time,omitempty"`

    // 重复信息
    PetInfo   string `json:"pet_info,omitempty"`

    // 图片信息
    PicUrls   string `json:"pic_urls,omitempty"`

    // 价格
    Price   int64 `json:"price,omitempty"`

    // 省份
    Province   int64 `json:"province,omitempty"`

    // 评论数
    RateNumber   int64 `json:"rate_number,omitempty"`

    // 评分
    RateScore   string `json:"rate_score,omitempty"`

    // 销量
    Sell   int64 `json:"sell,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // 服务设施集合类
    ShotelPropertiesVo   *ShotelPropertiesSetVo `json:"shotel_properties_vo,omitempty"`

    // 标准房型列表
    SroomTypes   []SRoomTypeVo `json:"sroom_types,omitempty"`

    // 档次，0--无档次，2--二星及经济，3--舒适，4--高档，5--豪华
    Star   string `json:"star,omitempty"`

    // 上下架状态，0--下架，其他状态-下架
    Status   int64 `json:"status,omitempty"`

    // 电话
    Tel   string `json:"tel,omitempty"`

    // 酒店业务分类
    Type   string `json:"type,omitempty"`

    // 客房数
    Rooms   int64 `json:"rooms,omitempty"`

}
