package travel

// PontusTravelFreedomItemExt 
/* model for simplify = false
type PontusTravelFreedomItemExt struct {

    // 大景点结构
    
    ScenicInfos  *struct {
        PontusTravelItemScenicInfo  *PontusTravelItemScenicInfo `json:"pontus_travel_item_scenic_info,omitempty"`
    } `json:"scenic_infos,omitempty"`
    

    // 资源结构
    
    OtherInfos  struct {
        PontusTravelItemResourceInfo  []PontusTravelItemResourceInfo `json:"pontus_travel_item_resource_info,omitempty"`
    } `json:"other_infos,omitempty"`
    

    // 酒店结构
    
    HotelInfos  *struct {
        PontusTravelItemHotelInfo  *PontusTravelItemHotelInfo `json:"pontus_travel_item_hotel_info,omitempty"`
    } `json:"hotel_infos,omitempty"`
    

    // 大交通结构
    
    BackTrafficInfo  *struct {
        PontusTravelItemTrafficInfo  *PontusTravelItemTrafficInfo `json:"pontus_travel_item_traffic_info,omitempty"`
    } `json:"back_traffic_info,omitempty"`
    

    // 自由行交通信息详细描述
    
    TrafficDesc   string `json:"traffic_desc,omitempty"`
    

    // 大交通结构
    
    GoTrafficInfo  *struct {
        PontusTravelItemTrafficInfo  *PontusTravelItemTrafficInfo `json:"pontus_travel_item_traffic_info,omitempty"`
    } `json:"go_traffic_info,omitempty"`
    

}
*/

// PontusTravelFreedomItemExt 
type PontusTravelFreedomItemExt struct {

    // 大景点结构
    ScenicInfos   *PontusTravelItemScenicInfo `json:"scenic_infos,omitempty"`

    // 资源结构
    OtherInfos   []PontusTravelItemResourceInfo `json:"other_infos,omitempty"`

    // 酒店结构
    HotelInfos   *PontusTravelItemHotelInfo `json:"hotel_infos,omitempty"`

    // 大交通结构
    BackTrafficInfo   *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty"`

    // 自由行交通信息详细描述
    TrafficDesc   string `json:"traffic_desc,omitempty"`

    // 大交通结构
    GoTrafficInfo   *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty"`

}
