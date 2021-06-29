package btrip

// HotelDto 
type HotelDto struct {
    // 地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 是否是客栈
    BnbHotel   bool `json:"bnb_hotel,omitempty" xml:"bnb_hotel,omitempty"`
    // 品牌
    Brand   string `json:"brand,omitempty" xml:"brand,omitempty"`
    // 入住时间
    CheckInTime   string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
    // 离店时间
    CheckOutTime   string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
    // 市级代码
    City   int64 `json:"city,omitempty" xml:"city,omitempty"`
    // 装修时间
    DecorateTime   string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
    // 酒店描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 地区代码
    District   int64 `json:"district,omitempty" xml:"district,omitempty"`
    // h5的detail页面的url
    H5DetailUrl   string `json:"h5_detail_url,omitempty" xml:"h5_detail_url,omitempty"`
    // 酒店设施
    HotelFacilities   string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty"`
    // 维度
    Lat   string `json:"lat,omitempty" xml:"lat,omitempty"`
    // 经度
    Lng   string `json:"lng,omitempty" xml:"lng,omitempty"`
    // 酒店名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 开业时间
    OpeningTime   string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
    // 酒店详情页的URL
    PcDetailUrl   string `json:"pc_detail_url,omitempty" xml:"pc_detail_url,omitempty"`
    // 图片地址
    PicUrls   string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
    // 省级diamante
    Province   int64 `json:"province,omitempty" xml:"province,omitempty"`
    // 评论数
    RateNumber   int64 `json:"rate_number,omitempty" xml:"rate_number,omitempty"`
    // 评分
    RateScore   string `json:"rate_score,omitempty" xml:"rate_score,omitempty"`
    // 房间列表
    Rooms   []RoomTypeDto `json:"rooms,omitempty" xml:"rooms>room_type_dto,omitempty"`
    // 服务设施
    ServicesStr   string `json:"services_str,omitempty" xml:"services_str,omitempty"`
    // 酒店标准ID
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    // 星级
    Star   string `json:"star,omitempty" xml:"star,omitempty"`
    // 状态，0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；默认为0
    Status   *model.File `json:"status,omitempty" xml:"status,omitempty"`
    // 供应商ID
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    // 电话
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    // 酒店类型
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
