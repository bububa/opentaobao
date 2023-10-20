package campus

import (
	"sync"
)

// FloorQuery 结构体
type FloorQuery struct {
	// 楼宇ID
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
}

var poolFloorQuery = sync.Pool{
	New: func() any {
		return new(FloorQuery)
	},
}

// GetFloorQuery() 从对象池中获取FloorQuery
func GetFloorQuery() *FloorQuery {
	return poolFloorQuery.Get().(*FloorQuery)
}

// ReleaseFloorQuery 释放FloorQuery
func ReleaseFloorQuery(v *FloorQuery) {
	v.BuildingId = 0
	poolFloorQuery.Put(v)
}
