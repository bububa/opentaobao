package campus

import (
	"sync"
)

// SpaceIdsQuery 结构体
type SpaceIdsQuery struct {
	// ids
	Ids []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
	// building/floor
	SpaceType string `json:"space_type,omitempty" xml:"space_type,omitempty"`
}

var poolSpaceIdsQuery = sync.Pool{
	New: func() any {
		return new(SpaceIdsQuery)
	},
}

// GetSpaceIdsQuery() 从对象池中获取SpaceIdsQuery
func GetSpaceIdsQuery() *SpaceIdsQuery {
	return poolSpaceIdsQuery.Get().(*SpaceIdsQuery)
}

// ReleaseSpaceIdsQuery 释放SpaceIdsQuery
func ReleaseSpaceIdsQuery(v *SpaceIdsQuery) {
	v.Ids = v.Ids[:0]
	v.SpaceType = ""
	poolSpaceIdsQuery.Put(v)
}
