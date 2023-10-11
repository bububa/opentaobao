package hotel

// HotelPrice 结构体
type HotelPrice struct {
	// 房型报价列表
	RoomPrices []RoomPrice `json:"room_prices,omitempty" xml:"room_prices>room_price,omitempty"`
	// 酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
