package alihouse

import (
	"sync"
)

// LinkInfoReqDto 结构体
type LinkInfoReqDto struct {
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolLinkInfoReqDto = sync.Pool{
	New: func() any {
		return new(LinkInfoReqDto)
	},
}

// GetLinkInfoReqDto() 从对象池中获取LinkInfoReqDto
func GetLinkInfoReqDto() *LinkInfoReqDto {
	return poolLinkInfoReqDto.Get().(*LinkInfoReqDto)
}

// ReleaseLinkInfoReqDto 释放LinkInfoReqDto
func ReleaseLinkInfoReqDto(v *LinkInfoReqDto) {
	v.Type = 0
	poolLinkInfoReqDto.Put(v)
}
