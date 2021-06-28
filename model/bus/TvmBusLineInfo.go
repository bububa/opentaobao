package bus

// TvmBusLineInfo 
/* model for simplify = false
type TvmBusLineInfo struct {

    // 出发时间 yyyy-mm-dd HH:mm:ss
    
    DepTime   string `json:"dep_time,omitempty"`
    

    // 距离 km
    
    Distance   int64 `json:"distance,omitempty"`
    

    // 到达目的地
    
    LastPlaceName   string `json:"last_place_name,omitempty"`
    

    // 出发城市(必须填写，后续程序中会进行校验)
    
    StartCityName   string `json:"start_city_name,omitempty"`
    

    // 出发省份(必须填写，后续程序中会进行强制校验，参考标准区域码表)
    
    StartProvinceName   string `json:"start_province_name,omitempty"`
    

    // 出发车站ID(必须填写，且与线上车站id能进行数据互通)
    
    StartStationId   string `json:"start_station_id,omitempty"`
    

    // 出发车站名称(必须填写，与线上售卖车站名称保持一致)
    
    StartStationName   string `json:"start_station_name,omitempty"`
    

    // 到达车站城市名称
    
    ToStationCityName   string `json:"to_station_city_name,omitempty"`
    

    // 到达车站ID
    
    ToStationId   string `json:"to_station_id,omitempty"`
    

    // 到达车站名称
    
    ToStationName   string `json:"to_station_name,omitempty"`
    

    // 到达车站省份
    
    ToStationProvinceName   string `json:"to_station_province_name,omitempty"`
    

    // 终点站名称
    
    Terminal   string `json:"terminal,omitempty"`
    

    // 车次编号
    
    BusNumber   string `json:"bus_number,omitempty"`
    

    // 车站地址(必须填写，后续程序中会进行强制校验，参考标准区域码表)
    
    StartStationAddress   string `json:"start_station_address,omitempty"`
    

    // 车型
    
    BusType   string `json:"bus_type,omitempty"`
    

    // 运行时长（分）
    
    Runtime   int64 `json:"runtime,omitempty"`
    

    // 出发省份码
    
    StartProvinceCode   string `json:"start_province_code,omitempty"`
    

    // 出发城市码
    
    StartCityCode   string `json:"start_city_code,omitempty"`
    

    // 到达省份码
    
    ToStationProvinceCode   string `json:"to_station_province_code,omitempty"`
    

    // 到达城市码
    
    ToStationCityCode   string `json:"to_station_city_code,omitempty"`
    

}
*/

// TvmBusLineInfo 
type TvmBusLineInfo struct {

    // 出发时间 yyyy-mm-dd HH:mm:ss
    DepTime   string `json:"dep_time,omitempty"`

    // 距离 km
    Distance   int64 `json:"distance,omitempty"`

    // 到达目的地
    LastPlaceName   string `json:"last_place_name,omitempty"`

    // 出发城市(必须填写，后续程序中会进行校验)
    StartCityName   string `json:"start_city_name,omitempty"`

    // 出发省份(必须填写，后续程序中会进行强制校验，参考标准区域码表)
    StartProvinceName   string `json:"start_province_name,omitempty"`

    // 出发车站ID(必须填写，且与线上车站id能进行数据互通)
    StartStationId   string `json:"start_station_id,omitempty"`

    // 出发车站名称(必须填写，与线上售卖车站名称保持一致)
    StartStationName   string `json:"start_station_name,omitempty"`

    // 到达车站城市名称
    ToStationCityName   string `json:"to_station_city_name,omitempty"`

    // 到达车站ID
    ToStationId   string `json:"to_station_id,omitempty"`

    // 到达车站名称
    ToStationName   string `json:"to_station_name,omitempty"`

    // 到达车站省份
    ToStationProvinceName   string `json:"to_station_province_name,omitempty"`

    // 终点站名称
    Terminal   string `json:"terminal,omitempty"`

    // 车次编号
    BusNumber   string `json:"bus_number,omitempty"`

    // 车站地址(必须填写，后续程序中会进行强制校验，参考标准区域码表)
    StartStationAddress   string `json:"start_station_address,omitempty"`

    // 车型
    BusType   string `json:"bus_type,omitempty"`

    // 运行时长（分）
    Runtime   int64 `json:"runtime,omitempty"`

    // 出发省份码
    StartProvinceCode   string `json:"start_province_code,omitempty"`

    // 出发城市码
    StartCityCode   string `json:"start_city_code,omitempty"`

    // 到达省份码
    ToStationProvinceCode   string `json:"to_station_province_code,omitempty"`

    // 到达城市码
    ToStationCityCode   string `json:"to_station_city_code,omitempty"`

}
