package axintrade

// Board 结构体
type Board struct {
	// 餐食类型
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
	// 餐食份数
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
}
