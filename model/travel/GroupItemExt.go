package travel

// GroupItemExt 
/* model for simplify = false
type GroupItemExt struct {

    // 必填，回程交通信息
    
    BackTrafficInfo  *struct {
        ItemTrafficInfo  *ItemTrafficInfo `json:"item_traffic_info,omitempty"`
    } `json:"back_traffic_info,omitempty"`
    

    // 是否支持电子合同，默认不支持
    
    Electronic   bool `json:"electronic,omitempty"`
    

    // 集合地信息
    
    GatherPlaces  *struct {
        GatherPlaceInfo  *GatherPlaceInfo `json:"gather_place_info,omitempty"`
    } `json:"gather_places,omitempty"`
    

    // 必填，去程交通信息
    
    GoTrafficInfo  *struct {
        ItemTrafficInfo  *ItemTrafficInfo `json:"item_traffic_info,omitempty"`
    } `json:"go_traffic_info,omitempty"`
    

    // 必填，线路类型，0 为目的地参团 	1为出发地参团
    
    RouteType   int64 `json:"route_type,omitempty"`
    

}
*/

// GroupItemExt 
type GroupItemExt struct {

    // 必填，回程交通信息
    BackTrafficInfo   *ItemTrafficInfo `json:"back_traffic_info,omitempty"`

    // 是否支持电子合同，默认不支持
    Electronic   bool `json:"electronic,omitempty"`

    // 集合地信息
    GatherPlaces   *GatherPlaceInfo `json:"gather_places,omitempty"`

    // 必填，去程交通信息
    GoTrafficInfo   *ItemTrafficInfo `json:"go_traffic_info,omitempty"`

    // 必填，线路类型，0 为目的地参团 	1为出发地参团
    RouteType   int64 `json:"route_type,omitempty"`

}
