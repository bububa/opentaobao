package tanx

import (
	"sync"
)

// DicDto 结构体
type DicDto struct {
	// 数据项值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 数据项ID
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolDicDto = sync.Pool{
	New: func() any {
		return new(DicDto)
	},
}

// GetDicDto() 从对象池中获取DicDto
func GetDicDto() *DicDto {
	return poolDicDto.Get().(*DicDto)
}

// ReleaseDicDto 释放DicDto
func ReleaseDicDto(v *DicDto) {
	v.Value = ""
	v.Code = 0
	poolDicDto.Put(v)
}
