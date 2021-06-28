package travel

// FreedomItemExt 
type FreedomItemExt struct {

    // 回程交通信息
    
    BackTrafficInfo   *ItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
    

    // 去程交通信息
    
    GoTrafficInfo   *ItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
    

    // 酒店信息
    
    HotelInfos   *ItemHotelInfo `json:"hotel_infos,omitempty" xml:"hotel_infos,omitempty"`
    

    // 其他资源信息
    
    OtherInfos   []ItemResourceInfo `json:"other_infos,omitempty" xml:"other_infos,omitempty"`
    

    // 景点信息
    
    ScenicInfos   *ItemScenicInfo `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
    

    // 自由行交通信息详细描述
    
    TrafficDesc   string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
    

}
