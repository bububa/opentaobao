package tmallservice

import (
	"sync"
)

// MemoDto 结构体
type MemoDto struct {
	// 备注内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 备注创建时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}

var poolMemoDto = sync.Pool{
	New: func() any {
		return new(MemoDto)
	},
}

// GetMemoDto() 从对象池中获取MemoDto
func GetMemoDto() *MemoDto {
	return poolMemoDto.Get().(*MemoDto)
}

// ReleaseMemoDto 释放MemoDto
func ReleaseMemoDto(v *MemoDto) {
	v.Content = ""
	v.Time = ""
	poolMemoDto.Put(v)
}
