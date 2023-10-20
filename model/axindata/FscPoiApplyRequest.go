package axindata

import (
	"sync"
)

// FscPoiApplyRequest 结构体
type FscPoiApplyRequest struct {
	// POI新增列表
	PoiList []FscPoiApplyApiDto `json:"poi_list,omitempty" xml:"poi_list>fsc_poi_apply_api_dto,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscPoiApplyRequest = sync.Pool{
	New: func() any {
		return new(FscPoiApplyRequest)
	},
}

// GetFscPoiApplyRequest() 从对象池中获取FscPoiApplyRequest
func GetFscPoiApplyRequest() *FscPoiApplyRequest {
	return poolFscPoiApplyRequest.Get().(*FscPoiApplyRequest)
}

// ReleaseFscPoiApplyRequest 释放FscPoiApplyRequest
func ReleaseFscPoiApplyRequest(v *FscPoiApplyRequest) {
	v.PoiList = v.PoiList[:0]
	v.SupplierId = ""
	poolFscPoiApplyRequest.Put(v)
}
