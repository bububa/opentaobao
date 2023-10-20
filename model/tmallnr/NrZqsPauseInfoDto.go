package tmallnr

import (
	"sync"
)

// NrZqsPauseInfoDto 结构体
type NrZqsPauseInfoDto struct {
	// 暂停开始时间，包含该天
	PauseStartDay string `json:"pause_start_day,omitempty" xml:"pause_start_day,omitempty"`
	// 暂停结束时间，包含该天
	PauseEndDay string `json:"pause_end_day,omitempty" xml:"pause_end_day,omitempty"`
}

var poolNrZqsPauseInfoDto = sync.Pool{
	New: func() any {
		return new(NrZqsPauseInfoDto)
	},
}

// GetNrZqsPauseInfoDto() 从对象池中获取NrZqsPauseInfoDto
func GetNrZqsPauseInfoDto() *NrZqsPauseInfoDto {
	return poolNrZqsPauseInfoDto.Get().(*NrZqsPauseInfoDto)
}

// ReleaseNrZqsPauseInfoDto 释放NrZqsPauseInfoDto
func ReleaseNrZqsPauseInfoDto(v *NrZqsPauseInfoDto) {
	v.PauseStartDay = ""
	v.PauseEndDay = ""
	poolNrZqsPauseInfoDto.Put(v)
}
