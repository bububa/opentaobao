package travel

// PontusTravelGroupItemExt 结构体
type PontusTravelGroupItemExt struct {
	// 必填，线路类型，0 为目的地参团 	1为出发地参团
	RouteType int64 `json:"route_type,omitempty" xml:"route_type,omitempty"`
	// 集合地结构
	GatherPlaces []PontusTravelGatherPlaceInfo `json:"gather_places,omitempty" xml:"gather_places>pontus_travel_gather_place_info,omitempty"`
	// 必填，回程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
	BackTrafficInfo *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 必填，去程交通。注意：目前跟团游交通只会存储交通方式，结构其他字段暂时不支持
	GoTrafficInfo *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
	// 是否支持电子合同，默认不支持
	Electronic bool `json:"electronic,omitempty" xml:"electronic,omitempty"`
	// 是否“纯玩团”，默认不传则是“含购物团”
	PurePlay bool `json:"pure_play,omitempty" xml:"pure_play,omitempty"`
}
