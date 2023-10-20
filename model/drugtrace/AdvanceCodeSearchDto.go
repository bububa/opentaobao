package drugtrace

import (
	"sync"
)

// AdvanceCodeSearchDto 结构体
type AdvanceCodeSearchDto struct {
	// 顶层结构
	PiatsCodeFlowResponseDTO *PiatsCodeFlowResponseDto `json:"piats_code_flow_response_d_t_o,omitempty" xml:"piats_code_flow_response_d_t_o,omitempty"`
}

var poolAdvanceCodeSearchDto = sync.Pool{
	New: func() any {
		return new(AdvanceCodeSearchDto)
	},
}

// GetAdvanceCodeSearchDto() 从对象池中获取AdvanceCodeSearchDto
func GetAdvanceCodeSearchDto() *AdvanceCodeSearchDto {
	return poolAdvanceCodeSearchDto.Get().(*AdvanceCodeSearchDto)
}

// ReleaseAdvanceCodeSearchDto 释放AdvanceCodeSearchDto
func ReleaseAdvanceCodeSearchDto(v *AdvanceCodeSearchDto) {
	v.PiatsCodeFlowResponseDTO = nil
	poolAdvanceCodeSearchDto.Put(v)
}
