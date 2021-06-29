package btrip

// StationLatitudeLongitudeVo 
type StationLatitudeLongitudeVo struct {
    // 车站经度
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    // 车站纬度
    Longitude   string `json:"longitude,omitempty" xml:"longitude,omitempty"`
    // 目的地频道提供的景点图片url（会有多张图片）
    PoiPictureUrlList   []string `json:"poi_picture_url_list,omitempty" xml:"poi_picture_url_list>string,omitempty"`
    // 目的地频道提供的标签(可能会有多个标签)
    PoiTagList   []string `json:"poi_tag_list,omitempty" xml:"poi_tag_list>string,omitempty"`
    // 车站地址
    StationAddress   string `json:"station_address,omitempty" xml:"station_address,omitempty"`
    // 车站名
    StationName   string `json:"station_name,omitempty" xml:"station_name,omitempty"`
    // 车站电话
    StationTel   string `json:"station_tel,omitempty" xml:"station_tel,omitempty"`
    // 00去程上车点，01去程下车点，10返程上车点，11返程下车点
    TourStationType   string `json:"tour_station_type,omitempty" xml:"tour_station_type,omitempty"`
}
