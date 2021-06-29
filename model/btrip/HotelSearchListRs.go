package btrip

// HotelSearchListRs 
type HotelSearchListRs struct {
    // 酒店列表
    Hotels   []HotelListDto `json:"hotels,omitempty" xml:"hotels>hotel_list_dto,omitempty"`
    // 酒店数量
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
