package bus

// BusNumberSearchRq 
/* model for simplify = false
type BusNumberSearchRq struct {

    // 出发城市
    
    DepCityName   string `json:"dep_city_name,omitempty"`
    

    // 出入日期
    
    DepDate   string `json:"dep_date,omitempty"`
    

    // 目的地
    
    LastPlaceName   string `json:"last_place_name,omitempty"`
    

}
*/

// BusNumberSearchRq 
type BusNumberSearchRq struct {

    // 出发城市
    DepCityName   string `json:"dep_city_name,omitempty"`

    // 出入日期
    DepDate   string `json:"dep_date,omitempty"`

    // 目的地
    LastPlaceName   string `json:"last_place_name,omitempty"`

}
