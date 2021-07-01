package btrip

// RoomInfoDo 结构体
type RoomInfoDo struct {
	// 房间名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 房间价格（分）
	RoomPrice int64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	// 房间数量
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// 总间夜
	NightNumber int64 `json:"night_number,omitempty" xml:"night_number,omitempty"`
	// 货币种类
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
}
