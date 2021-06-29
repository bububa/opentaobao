package btrip

// CarInfoDO 
type CarInfoDO struct {
    // 预订出发地
    FromAddr   string `json:"from_addr,omitempty" xml:"from_addr,omitempty"`
    // 预订到达地
    ToAddr   string `json:"to_addr,omitempty" xml:"to_addr,omitempty"`
    // 预订出发城市
    FromCityName   string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
    // 预订到达城市
    ToCityName   string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
    // 实际出发城市
    RealFromCityName   string `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
    // 实际到达城市
    RealToCityName   string `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
    // 2:滴滴出行 8:滴滴出行 100087:滴滴出行 100000:滴滴出行 100003:曹操出行 3:曹操出行 100085:曹操出行 100007:阳光出行 5:阳光出行 0：其他
    Provider   int64 `json:"provider,omitempty" xml:"provider,omitempty"`
    // 滴滴出行 曹操出行 阳光出行 其他
    ProviderName   string `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
    // 打车是由
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // TAXI:出租车 KC:经济型 COSY:舒适型 COMMERCE:商务型 LUXURY:豪华型
    CarLevel   string `json:"car_level,omitempty" xml:"car_level,omitempty"`
    // 车辆类型
    CarInfo   string `json:"car_info,omitempty" xml:"car_info,omitempty"`
    // 乘客发布用车时间
    PublishTime   string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
    // 乘客上车时间
    TakenTime   string `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
    // 司机到达确认时间
    DriverConfirmTime   string `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
    // 行驶公里数（km）
    TravelDistance   string `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
    // 实际出发地
    RealFromAddr   string `json:"real_from_addr,omitempty" xml:"real_from_addr,omitempty"`
    // 实际到达地
    RealToAddr   string `json:"real_to_addr,omitempty" xml:"real_to_addr,omitempty"`
}
