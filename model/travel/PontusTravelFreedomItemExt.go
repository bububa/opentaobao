package travel

// PontusTravelFreedomItemExt 结构体
type PontusTravelFreedomItemExt struct {
	// 大景点结构
	ScenicInfos *PontusTravelItemScenicInfo `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
	// 资源结构
	OtherInfos []PontusTravelItemResourceInfo `json:"other_infos,omitempty" xml:"other_infos>pontus_travel_item_resource_info,omitempty"`
	// 酒店结构
	HotelInfos *PontusTravelItemHotelInfo `json:"hotel_infos,omitempty" xml:"hotel_infos,omitempty"`
	// 大交通结构
	BackTrafficInfo *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 自由行交通信息详细描述
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 大交通结构
	GoTrafficInfo *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
}
