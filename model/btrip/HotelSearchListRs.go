package btrip

// HotelSearchListRS 
type HotelSearchListRS struct {
    // 酒店列表
    Hotels   []HotelListDTO `json:"hotels,omitempty" xml:"hotels>hotel_list_dto,omitempty"`
    // 酒店数量
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
