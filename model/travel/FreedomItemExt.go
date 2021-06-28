package travel

// FreedomItemExt 
/* model for simplify = false
type FreedomItemExt struct {

    // 回程交通信息
    
    BackTrafficInfo  *struct {
        ItemTrafficInfo  *ItemTrafficInfo `json:"item_traffic_info,omitempty"`
    } `json:"back_traffic_info,omitempty"`
    

    // 去程交通信息
    
    GoTrafficInfo  *struct {
        ItemTrafficInfo  *ItemTrafficInfo `json:"item_traffic_info,omitempty"`
    } `json:"go_traffic_info,omitempty"`
    

    // 酒店信息
    
    HotelInfos  *struct {
        ItemHotelInfo  *ItemHotelInfo `json:"item_hotel_info,omitempty"`
    } `json:"hotel_infos,omitempty"`
    

    // 其他资源信息
    
    OtherInfos  struct {
        ItemResourceInfo  []ItemResourceInfo `json:"item_resource_info,omitempty"`
    } `json:"other_infos,omitempty"`
    

    // 景点信息
    
    ScenicInfos  *struct {
        ItemScenicInfo  *ItemScenicInfo `json:"item_scenic_info,omitempty"`
    } `json:"scenic_infos,omitempty"`
    

    // 自由行交通信息详细描述
    
    TrafficDesc   string `json:"traffic_desc,omitempty"`
    

}
*/

// FreedomItemExt 
type FreedomItemExt struct {

    // 回程交通信息
    BackTrafficInfo   *ItemTrafficInfo `json:"back_traffic_info,omitempty"`

    // 去程交通信息
    GoTrafficInfo   *ItemTrafficInfo `json:"go_traffic_info,omitempty"`

    // 酒店信息
    HotelInfos   *ItemHotelInfo `json:"hotel_infos,omitempty"`

    // 其他资源信息
    OtherInfos   []ItemResourceInfo `json:"other_infos,omitempty"`

    // 景点信息
    ScenicInfos   *ItemScenicInfo `json:"scenic_infos,omitempty"`

    // 自由行交通信息详细描述
    TrafficDesc   string `json:"traffic_desc,omitempty"`

}
