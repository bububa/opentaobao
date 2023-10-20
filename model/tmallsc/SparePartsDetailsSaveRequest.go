package tmallsc

import (
	"sync"
)

// SparePartsDetailsSaveRequest 结构体
type SparePartsDetailsSaveRequest struct {
	// 备件申请单信息
	ApplicationInfoDto *ApplicationInfoDto `json:"application_info_dto,omitempty" xml:"application_info_dto,omitempty"`
	// 备件明细
	SparePartsInfoDto *SparePartsInfoDto `json:"spare_parts_info_dto,omitempty" xml:"spare_parts_info_dto,omitempty"`
	// 天猫服务工单号
	WorkCardId int64 `json:"work_card_id,omitempty" xml:"work_card_id,omitempty"`
}

var poolSparePartsDetailsSaveRequest = sync.Pool{
	New: func() any {
		return new(SparePartsDetailsSaveRequest)
	},
}

// GetSparePartsDetailsSaveRequest() 从对象池中获取SparePartsDetailsSaveRequest
func GetSparePartsDetailsSaveRequest() *SparePartsDetailsSaveRequest {
	return poolSparePartsDetailsSaveRequest.Get().(*SparePartsDetailsSaveRequest)
}

// ReleaseSparePartsDetailsSaveRequest 释放SparePartsDetailsSaveRequest
func ReleaseSparePartsDetailsSaveRequest(v *SparePartsDetailsSaveRequest) {
	v.ApplicationInfoDto = nil
	v.SparePartsInfoDto = nil
	v.WorkCardId = 0
	poolSparePartsDetailsSaveRequest.Put(v)
}
