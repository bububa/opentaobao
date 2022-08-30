package btrip

// BtripHotelBoardDto 结构体
type BtripHotelBoardDto struct {
	// 餐食数量
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
	// 餐食类型
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
}
