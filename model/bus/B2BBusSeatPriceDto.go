package bus

// B2BBusSeatPriceDto 
type B2BBusSeatPriceDto struct {

    // 出发城市
    DepCityName   string `json:"dep_city_name,omitempty"`

    // 出发时间
    DepTime   string `json:"dep_time,omitempty"`

    // 车次全价
    FullPrice   int64 `json:"full_price,omitempty"`

    // 目的地
    LastPlaceName   string `json:"last_place_name,omitempty"`

    // 车次ID
    ScheduleId   string `json:"schedule_id,omitempty"`

    // 余票数量
    Stock   int64 `json:"stock,omitempty"`

    // 服务费
    ServicePrice   int64 `json:"service_price,omitempty"`

}
