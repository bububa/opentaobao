package bus

// TopBusNumerPushRq 
/* model for simplify = false
type TopBusNumerPushRq struct {

    // 目的地
    
    LastPlaceName   string `json:"last_place_name,omitempty"`
    

    // 出发日期
    
    DepDate   string `json:"dep_date,omitempty"`
    

    // 出发城市
    
    FromCityName   string `json:"from_city_name,omitempty"`
    

}
*/

// TopBusNumerPushRq 
type TopBusNumerPushRq struct {

    // 目的地
    LastPlaceName   string `json:"last_place_name,omitempty"`

    // 出发日期
    DepDate   string `json:"dep_date,omitempty"`

    // 出发城市
    FromCityName   string `json:"from_city_name,omitempty"`

}
