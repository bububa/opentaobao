package cainiaohandover

import (
	"sync"
)

// OpenHandoverDraftSaveResponse 结构体
type OpenHandoverDraftSaveResponse struct {
	// 交接批次号，即交接单id
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
	// 交接物id，即大包id
	HandoverContentId int64 `json:"handover_content_id,omitempty" xml:"handover_content_id,omitempty"`
}

var poolOpenHandoverDraftSaveResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverDraftSaveResponse)
	},
}

// GetOpenHandoverDraftSaveResponse() 从对象池中获取OpenHandoverDraftSaveResponse
func GetOpenHandoverDraftSaveResponse() *OpenHandoverDraftSaveResponse {
	return poolOpenHandoverDraftSaveResponse.Get().(*OpenHandoverDraftSaveResponse)
}

// ReleaseOpenHandoverDraftSaveResponse 释放OpenHandoverDraftSaveResponse
func ReleaseOpenHandoverDraftSaveResponse(v *OpenHandoverDraftSaveResponse) {
	v.HandoverOrderId = 0
	v.HandoverContentId = 0
	poolOpenHandoverDraftSaveResponse.Put(v)
}
