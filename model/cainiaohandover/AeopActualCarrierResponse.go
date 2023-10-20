package cainiaohandover

import (
	"sync"
)

// AeopActualCarrierResponse 结构体
type AeopActualCarrierResponse struct {
	// 实际承运商
	CourierList []Courierlist `json:"courier_list,omitempty" xml:"courier_list>courierlist,omitempty"`
}

var poolAeopActualCarrierResponse = sync.Pool{
	New: func() any {
		return new(AeopActualCarrierResponse)
	},
}

// GetAeopActualCarrierResponse() 从对象池中获取AeopActualCarrierResponse
func GetAeopActualCarrierResponse() *AeopActualCarrierResponse {
	return poolAeopActualCarrierResponse.Get().(*AeopActualCarrierResponse)
}

// ReleaseAeopActualCarrierResponse 释放AeopActualCarrierResponse
func ReleaseAeopActualCarrierResponse(v *AeopActualCarrierResponse) {
	v.CourierList = v.CourierList[:0]
	poolAeopActualCarrierResponse.Put(v)
}
