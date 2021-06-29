package cainiaolocker

// StationInfo 
type StationInfo struct {

    // 邮编
    
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    

    // 0-上线，1-下线
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 小区
    
    HousingEstate   string `json:"housing_estate,omitempty" xml:"housing_estate,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 站点照片url
    
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    

    // 省份
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 坐标类型，MARS-火星坐标（高德坐标），BAIDU-百度坐标,GPS-GPS坐标
    
    CoordType   string `json:"coord_type,omitempty" xml:"coord_type,omitempty"`
    

    // 街道
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 区县
    
    District   string `json:"district,omitempty" xml:"district,omitempty"`
    

    // 站点服务能力描述
    
    Extra   string `json:"extra,omitempty" xml:"extra,omitempty"`
    

    // 站点类型：100-代收点
    
    StationType   int64 `json:"station_type,omitempty" xml:"station_type,omitempty"`
    

    // 站点联系方式
    
    Contact   string `json:"contact,omitempty" xml:"contact,omitempty"`
    

    // 站点经度
    
    StationLng   string `json:"station_lng,omitempty" xml:"station_lng,omitempty"`
    

    // 站点详细地址
    
    StationAddr   string `json:"station_addr,omitempty" xml:"station_addr,omitempty"`
    

    // 站点id
    
    StationId   string `json:"station_id,omitempty" xml:"station_id,omitempty"`
    

    // 站点名字
    
    StationName   string `json:"station_name,omitempty" xml:"station_name,omitempty"`
    

    // 站点编码
    
    StationNo   string `json:"station_no,omitempty" xml:"station_no,omitempty"`
    

    // 站点纬度
    
    StationLat   string `json:"station_lat,omitempty" xml:"station_lat,omitempty"`
    

}
