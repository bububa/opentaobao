package alitrippoi

import (
	"sync"
)

// FliggyPoiIdParam 结构体
type FliggyPoiIdParam struct {
	// 需要查询的poiid
	PoiIds []string `json:"poi_ids,omitempty" xml:"poi_ids>string,omitempty"`
}

var poolFliggyPoiIdParam = sync.Pool{
	New: func() any {
		return new(FliggyPoiIdParam)
	},
}

// GetFliggyPoiIdParam() 从对象池中获取FliggyPoiIdParam
func GetFliggyPoiIdParam() *FliggyPoiIdParam {
	return poolFliggyPoiIdParam.Get().(*FliggyPoiIdParam)
}

// ReleaseFliggyPoiIdParam 释放FliggyPoiIdParam
func ReleaseFliggyPoiIdParam(v *FliggyPoiIdParam) {
	v.PoiIds = v.PoiIds[:0]
	poolFliggyPoiIdParam.Put(v)
}
