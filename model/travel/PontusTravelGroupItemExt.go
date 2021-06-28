package travel

// PontusTravelGroupItemExt 
/* model for simplify = false
type PontusTravelGroupItemExt struct {

    // 必填，线路类型，0 为目的地参团 	1为出发地参团
    
    RouteType   int64 `json:"route_type,omitempty"`
    

    // 集合地结构
    
    GatherPlaces  struct {
        PontusTravelGatherPlaceInfo  []PontusTravelGatherPlaceInfo `json:"pontus_travel_gather_place_info,omitempty"`
    } `json:"gather_places,omitempty"`
    

    // 必填，回程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
    
    BackTrafficInfo  *struct {
        PontusTravelItemTrafficInfo  *PontusTravelItemTrafficInfo `json:"pontus_travel_item_traffic_info,omitempty"`
    } `json:"back_traffic_info,omitempty"`
    

    // 必填，去程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
    
    GoTrafficInfo  *struct {
        PontusTravelItemTrafficInfo  *PontusTravelItemTrafficInfo `json:"pontus_travel_item_traffic_info,omitempty"`
    } `json:"go_traffic_info,omitempty"`
    

    // 是否支持电子合同，默认不支持
    
    Electronic   bool `json:"electronic,omitempty"`
    

    // 是否“纯玩团”，默认不传则是“含购物团”
    
    PurePlay   bool `json:"pure_play,omitempty"`
    

}
*/

// PontusTravelGroupItemExt 
type PontusTravelGroupItemExt struct {

    // 必填，线路类型，0 为目的地参团 	1为出发地参团
    RouteType   int64 `json:"route_type,omitempty"`

    // 集合地结构
    GatherPlaces   []PontusTravelGatherPlaceInfo `json:"gather_places,omitempty"`

    // 必填，回程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
    BackTrafficInfo   *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty"`

    // 必填，去程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
    GoTrafficInfo   *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty"`

    // 是否支持电子合同，默认不支持
    Electronic   bool `json:"electronic,omitempty"`

    // 是否“纯玩团”，默认不传则是“含购物团”
    PurePlay   bool `json:"pure_play,omitempty"`

}
