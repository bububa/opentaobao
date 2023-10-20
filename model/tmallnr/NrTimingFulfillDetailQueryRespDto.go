package tmallnr

import (
	"sync"
)

// NrTimingFulfillDetailQueryRespDto 结构体
type NrTimingFulfillDetailQueryRespDto struct {
	// 历史状态
	NrDeliveryBriefStatusDTOs []NrDeliveryBriefStatusDto `json:"nr_delivery_brief_status_d_t_os,omitempty" xml:"nr_delivery_brief_status_d_t_os>nr_delivery_brief_status_dto,omitempty"`
	// 当前状态
	NrDeliveryBriefStatusDTO *NrDeliveryBriefStatusDto `json:"nr_delivery_brief_status_d_t_o,omitempty" xml:"nr_delivery_brief_status_d_t_o,omitempty"`
}

var poolNrTimingFulfillDetailQueryRespDto = sync.Pool{
	New: func() any {
		return new(NrTimingFulfillDetailQueryRespDto)
	},
}

// GetNrTimingFulfillDetailQueryRespDto() 从对象池中获取NrTimingFulfillDetailQueryRespDto
func GetNrTimingFulfillDetailQueryRespDto() *NrTimingFulfillDetailQueryRespDto {
	return poolNrTimingFulfillDetailQueryRespDto.Get().(*NrTimingFulfillDetailQueryRespDto)
}

// ReleaseNrTimingFulfillDetailQueryRespDto 释放NrTimingFulfillDetailQueryRespDto
func ReleaseNrTimingFulfillDetailQueryRespDto(v *NrTimingFulfillDetailQueryRespDto) {
	v.NrDeliveryBriefStatusDTOs = v.NrDeliveryBriefStatusDTOs[:0]
	v.NrDeliveryBriefStatusDTO = nil
	poolNrTimingFulfillDetailQueryRespDto.Put(v)
}
