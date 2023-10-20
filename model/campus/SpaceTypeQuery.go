package campus

import (
	"sync"
)

// SpaceTypeQuery 结构体
type SpaceTypeQuery struct {
	// 模糊查询key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 每页大小
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 大类id
	TopLevelId int64 `json:"top_level_id,omitempty" xml:"top_level_id,omitempty"`
	// 中类id
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// 分组大类id
	GroupTopLevelId int64 `json:"group_top_level_id,omitempty" xml:"group_top_level_id,omitempty"`
	// 小类id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 类别，室内，室外，逻辑
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 起始行
	StartRow int64 `json:"start_row,omitempty" xml:"start_row,omitempty"`
	// 空间大类id
	SpaceTopLevelId int64 `json:"space_top_level_id,omitempty" xml:"space_top_level_id,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolSpaceTypeQuery = sync.Pool{
	New: func() any {
		return new(SpaceTypeQuery)
	},
}

// GetSpaceTypeQuery() 从对象池中获取SpaceTypeQuery
func GetSpaceTypeQuery() *SpaceTypeQuery {
	return poolSpaceTypeQuery.Get().(*SpaceTypeQuery)
}

// ReleaseSpaceTypeQuery 释放SpaceTypeQuery
func ReleaseSpaceTypeQuery(v *SpaceTypeQuery) {
	v.Key = ""
	v.Limit = 0
	v.TopLevelId = 0
	v.Pid = 0
	v.GroupTopLevelId = 0
	v.Id = 0
	v.Category = 0
	v.StartRow = 0
	v.SpaceTopLevelId = 0
	v.CurrentPage = 0
	poolSpaceTypeQuery.Put(v)
}
