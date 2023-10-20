package user

import (
	"sync"
)

// TopDivisionRecordReqDto 结构体
type TopDivisionRecordReqDto struct {
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 开始时间(毫秒为单位)
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间(毫秒为单位)
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolTopDivisionRecordReqDto = sync.Pool{
	New: func() any {
		return new(TopDivisionRecordReqDto)
	},
}

// GetTopDivisionRecordReqDto() 从对象池中获取TopDivisionRecordReqDto
func GetTopDivisionRecordReqDto() *TopDivisionRecordReqDto {
	return poolTopDivisionRecordReqDto.Get().(*TopDivisionRecordReqDto)
}

// ReleaseTopDivisionRecordReqDto 释放TopDivisionRecordReqDto
func ReleaseTopDivisionRecordReqDto(v *TopDivisionRecordReqDto) {
	v.PageNo = 0
	v.PageSize = 0
	v.StartTime = 0
	v.EndTime = 0
	poolTopDivisionRecordReqDto.Put(v)
}
