package axintrade

import (
	"sync"
)

// Board 结构体
type Board struct {
	// 餐食类型
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
	// 餐食份数
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
}

var poolBoard = sync.Pool{
	New: func() any {
		return new(Board)
	},
}

// GetBoard() 从对象池中获取Board
func GetBoard() *Board {
	return poolBoard.Get().(*Board)
}

// ReleaseBoard 释放Board
func ReleaseBoard(v *Board) {
	v.BoardType = 0
	v.BoardNum = 0
	poolBoard.Put(v)
}
