package campus

import (
	"sync"
)

// BuildingQuery 结构体
type BuildingQuery struct {
	// 园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
}

var poolBuildingQuery = sync.Pool{
	New: func() any {
		return new(BuildingQuery)
	},
}

// GetBuildingQuery() 从对象池中获取BuildingQuery
func GetBuildingQuery() *BuildingQuery {
	return poolBuildingQuery.Get().(*BuildingQuery)
}

// ReleaseBuildingQuery 释放BuildingQuery
func ReleaseBuildingQuery(v *BuildingQuery) {
	v.CampusId = 0
	poolBuildingQuery.Put(v)
}
