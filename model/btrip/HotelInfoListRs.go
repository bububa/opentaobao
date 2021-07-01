package btrip

// HotelInfoListRS 
type HotelInfoListRS struct {
    // 基础酒店数据列表
    Hotels   []HotelDTO `json:"hotels,omitempty" xml:"hotels>hotel_dto,omitempty"`
    // 酒店数量
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
