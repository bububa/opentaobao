package cainiaohandover

import (
	"sync"
)

// HandoverContentAddSubbagsDto 结构体
type HandoverContentAddSubbagsDto struct {
	// 追加的大包LP
	SubbagHandoverContentCode string `json:"subbag_handover_content_code,omitempty" xml:"subbag_handover_content_code,omitempty"`
	// 交接单id
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
	// 追加的大包id
	SubbagHandoverContentId int64 `json:"subbag_handover_content_id,omitempty" xml:"subbag_handover_content_id,omitempty"`
}

var poolHandoverContentAddSubbagsDto = sync.Pool{
	New: func() any {
		return new(HandoverContentAddSubbagsDto)
	},
}

// GetHandoverContentAddSubbagsDto() 从对象池中获取HandoverContentAddSubbagsDto
func GetHandoverContentAddSubbagsDto() *HandoverContentAddSubbagsDto {
	return poolHandoverContentAddSubbagsDto.Get().(*HandoverContentAddSubbagsDto)
}

// ReleaseHandoverContentAddSubbagsDto 释放HandoverContentAddSubbagsDto
func ReleaseHandoverContentAddSubbagsDto(v *HandoverContentAddSubbagsDto) {
	v.SubbagHandoverContentCode = ""
	v.HandoverOrderId = 0
	v.SubbagHandoverContentId = 0
	poolHandoverContentAddSubbagsDto.Put(v)
}
