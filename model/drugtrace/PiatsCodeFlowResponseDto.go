package drugtrace

import (
	"sync"
)

// PiatsCodeFlowResponseDto 结构体
type PiatsCodeFlowResponseDto struct {
	// 头部结构
	Header *Header `json:"header,omitempty" xml:"header,omitempty"`
	// 内容体
	ResponseBody *ResponseBody `json:"response_body,omitempty" xml:"response_body,omitempty"`
}

var poolPiatsCodeFlowResponseDto = sync.Pool{
	New: func() any {
		return new(PiatsCodeFlowResponseDto)
	},
}

// GetPiatsCodeFlowResponseDto() 从对象池中获取PiatsCodeFlowResponseDto
func GetPiatsCodeFlowResponseDto() *PiatsCodeFlowResponseDto {
	return poolPiatsCodeFlowResponseDto.Get().(*PiatsCodeFlowResponseDto)
}

// ReleasePiatsCodeFlowResponseDto 释放PiatsCodeFlowResponseDto
func ReleasePiatsCodeFlowResponseDto(v *PiatsCodeFlowResponseDto) {
	v.Header = nil
	v.ResponseBody = nil
	poolPiatsCodeFlowResponseDto.Put(v)
}
