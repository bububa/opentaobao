package logistic

import (
	"sync"
)

// UpdateResendStatusDto 结构体
type UpdateResendStatusDto struct {
	// 描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// erp补发单
	ReissueId string `json:"reissue_id,omitempty" xml:"reissue_id,omitempty"`
	// 平台补发单唯一标识
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 补发单状态（-1=关闭，1=补发成功，2=部分成功）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 主订单
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolUpdateResendStatusDto = sync.Pool{
	New: func() any {
		return new(UpdateResendStatusDto)
	},
}

// GetUpdateResendStatusDto() 从对象池中获取UpdateResendStatusDto
func GetUpdateResendStatusDto() *UpdateResendStatusDto {
	return poolUpdateResendStatusDto.Get().(*UpdateResendStatusDto)
}

// ReleaseUpdateResendStatusDto 释放UpdateResendStatusDto
func ReleaseUpdateResendStatusDto(v *UpdateResendStatusDto) {
	v.Msg = ""
	v.ReissueId = ""
	v.SourceId = ""
	v.Status = 0
	v.Tid = 0
	poolUpdateResendStatusDto.Put(v)
}
