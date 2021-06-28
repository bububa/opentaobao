package bus

// B2BBusLineInfo 
/* model for simplify = false
type B2BBusLineInfo struct {

    // 出发时间
    
    DepTime   string `json:"dep_time,omitempty"`
    

    // 目的地
    
    LastPlaceName   string `json:"last_place_name,omitempty"`
    

    // 车次ID
    
    ScheduleId   string `json:"schedule_id,omitempty"`
    

    // 出发城市
    
    StartCityName   string `json:"start_city_name,omitempty"`
    

    // 出发车站
    
    StartStationName   string `json:"start_station_name,omitempty"`
    

    // 出发车站地址
    
    StartStationNameAddress   string `json:"start_station_name_address,omitempty"`
    

    // 到达车站
    
    ToStationName   string `json:"to_station_name,omitempty"`
    

}
*/

// B2BBusLineInfo 
type B2BBusLineInfo struct {

    // 出发时间
    DepTime   string `json:"dep_time,omitempty"`

    // 目的地
    LastPlaceName   string `json:"last_place_name,omitempty"`

    // 车次ID
    ScheduleId   string `json:"schedule_id,omitempty"`

    // 出发城市
    StartCityName   string `json:"start_city_name,omitempty"`

    // 出发车站
    StartStationName   string `json:"start_station_name,omitempty"`

    // 出发车站地址
    StartStationNameAddress   string `json:"start_station_name_address,omitempty"`

    // 到达车站
    ToStationName   string `json:"to_station_name,omitempty"`

}
