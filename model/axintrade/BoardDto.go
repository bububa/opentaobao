package axintrade

import (
	"sync"
)

// BoardDto 结构体
type BoardDto struct {
	// 餐食数量
	BoardNum int64 `json:"board_num,omitempty" xml:"board_num,omitempty"`
	// 餐食种类
	BoardType int64 `json:"board_type,omitempty" xml:"board_type,omitempty"`
}

var poolBoardDto = sync.Pool{
	New: func() any {
		return new(BoardDto)
	},
}

// GetBoardDto() 从对象池中获取BoardDto
func GetBoardDto() *BoardDto {
	return poolBoardDto.Get().(*BoardDto)
}

// ReleaseBoardDto 释放BoardDto
func ReleaseBoardDto(v *BoardDto) {
	v.BoardNum = 0
	v.BoardType = 0
	poolBoardDto.Put(v)
}
