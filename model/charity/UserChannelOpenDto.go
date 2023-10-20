package charity

import (
	"sync"
)

// UserChannelOpenDto 结构体
type UserChannelOpenDto struct {
	// 用户开放标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolUserChannelOpenDto = sync.Pool{
	New: func() any {
		return new(UserChannelOpenDto)
	},
}

// GetUserChannelOpenDto() 从对象池中获取UserChannelOpenDto
func GetUserChannelOpenDto() *UserChannelOpenDto {
	return poolUserChannelOpenDto.Get().(*UserChannelOpenDto)
}

// ReleaseUserChannelOpenDto 释放UserChannelOpenDto
func ReleaseUserChannelOpenDto(v *UserChannelOpenDto) {
	v.OpenId = ""
	poolUserChannelOpenDto.Put(v)
}
