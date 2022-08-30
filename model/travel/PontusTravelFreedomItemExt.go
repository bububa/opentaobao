package travel

// PontusTravelFreedomItemExt 结构体
type PontusTravelFreedomItemExt struct {
	// 其他资源信息
	OtherInfos []PontusTravelItemResourceInfo `json:"other_infos,omitempty" xml:"other_infos>pontus_travel_item_resource_info,omitempty"`
	// 自由行交通信息详细描述
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 景点信息
	ScenicInfos *PontusTravelItemScenicInfo `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
	// 酒店信息
	HotelInfos *PontusTravelItemHotelInfo `json:"hotel_infos,omitempty" xml:"hotel_infos,omitempty"`
	// 回程交通信息
	BackTrafficInfo *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 去程交通信息
	GoTrafficInfo *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
}
