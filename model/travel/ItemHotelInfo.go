package travel

// ItemHotelInfo 
/* model for simplify = false
type ItemHotelInfo struct {

    // 结构化酒店信息 酒店结构化信息和文本描述二选一
    
    HotelList  struct {
        ItemHotel  []ItemHotel `json:"item_hotel,omitempty"`
    } `json:"hotel_list,omitempty"`
    

}
*/

// ItemHotelInfo 
type ItemHotelInfo struct {

    // 结构化酒店信息 酒店结构化信息和文本描述二选一
    HotelList   []ItemHotel `json:"hotel_list,omitempty"`

}
