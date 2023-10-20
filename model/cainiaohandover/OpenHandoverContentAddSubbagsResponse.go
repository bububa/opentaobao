package cainiaohandover

import (
	"sync"
)

// OpenHandoverContentAddSubbagsResponse 结构体
type OpenHandoverContentAddSubbagsResponse struct {
	// 追加大包列表
	SubbagHandoverContentList []HandoverContentAddSubbagsDto `json:"subbag_handover_content_list,omitempty" xml:"subbag_handover_content_list>handover_content_add_subbags_dto,omitempty"`
}

var poolOpenHandoverContentAddSubbagsResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverContentAddSubbagsResponse)
	},
}

// GetOpenHandoverContentAddSubbagsResponse() 从对象池中获取OpenHandoverContentAddSubbagsResponse
func GetOpenHandoverContentAddSubbagsResponse() *OpenHandoverContentAddSubbagsResponse {
	return poolOpenHandoverContentAddSubbagsResponse.Get().(*OpenHandoverContentAddSubbagsResponse)
}

// ReleaseOpenHandoverContentAddSubbagsResponse 释放OpenHandoverContentAddSubbagsResponse
func ReleaseOpenHandoverContentAddSubbagsResponse(v *OpenHandoverContentAddSubbagsResponse) {
	v.SubbagHandoverContentList = v.SubbagHandoverContentList[:0]
	poolOpenHandoverContentAddSubbagsResponse.Put(v)
}
