package cainiaohandover

import (
	"sync"
)

// OpenHandoverContentUpdateResponse 结构体
type OpenHandoverContentUpdateResponse struct {
	// 报错小包列表
	UpdateErrorParcelOrderList []HandoverContentUpdateErrorParcelDto `json:"update_error_parcel_order_list,omitempty" xml:"update_error_parcel_order_list>handover_content_update_error_parcel_dto,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolOpenHandoverContentUpdateResponse = sync.Pool{
	New: func() any {
		return new(OpenHandoverContentUpdateResponse)
	},
}

// GetOpenHandoverContentUpdateResponse() 从对象池中获取OpenHandoverContentUpdateResponse
func GetOpenHandoverContentUpdateResponse() *OpenHandoverContentUpdateResponse {
	return poolOpenHandoverContentUpdateResponse.Get().(*OpenHandoverContentUpdateResponse)
}

// ReleaseOpenHandoverContentUpdateResponse 释放OpenHandoverContentUpdateResponse
func ReleaseOpenHandoverContentUpdateResponse(v *OpenHandoverContentUpdateResponse) {
	v.UpdateErrorParcelOrderList = v.UpdateErrorParcelOrderList[:0]
	v.Result = false
	poolOpenHandoverContentUpdateResponse.Put(v)
}
