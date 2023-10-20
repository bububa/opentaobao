package cainiaohandover

import (
	"sync"
)

// HandoverContentUpdateErrorParcelDto 结构体
type HandoverContentUpdateErrorParcelDto struct {
	// 小包LP号
	LpCode string `json:"lp_code,omitempty" xml:"lp_code,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolHandoverContentUpdateErrorParcelDto = sync.Pool{
	New: func() any {
		return new(HandoverContentUpdateErrorParcelDto)
	},
}

// GetHandoverContentUpdateErrorParcelDto() 从对象池中获取HandoverContentUpdateErrorParcelDto
func GetHandoverContentUpdateErrorParcelDto() *HandoverContentUpdateErrorParcelDto {
	return poolHandoverContentUpdateErrorParcelDto.Get().(*HandoverContentUpdateErrorParcelDto)
}

// ReleaseHandoverContentUpdateErrorParcelDto 释放HandoverContentUpdateErrorParcelDto
func ReleaseHandoverContentUpdateErrorParcelDto(v *HandoverContentUpdateErrorParcelDto) {
	v.LpCode = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	poolHandoverContentUpdateErrorParcelDto.Put(v)
}
