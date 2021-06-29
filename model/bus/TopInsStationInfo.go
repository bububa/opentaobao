package bus

// TopInsStationInfo 
type TopInsStationInfo struct {
    // 机具ID
    MachineId   string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
    // isv名称
    IsvName   string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
    // isv-id
    IsvId   string `json:"isv_id,omitempty" xml:"isv_id,omitempty"`
    // 出发车站站点ID
    StartStationId   string `json:"start_station_id,omitempty" xml:"start_station_id,omitempty"`
    // 城市代码，数据源来自行业基础数据字典
    CityCode   string `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // 省份代码，数据源来自行业基础数据字典
    ProvinceCode   string `json:"province_code,omitempty" xml:"province_code,omitempty"`
}
