package campus

import (
	"sync"
)

// SpaceGroupQuery 结构体
type SpaceGroupQuery struct {
	// 分组ID集合
	Ids []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
	// 类型编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// 空间分组编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 分组名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分组名称或者code
	NameOrCode string `json:"name_or_code,omitempty" xml:"name_or_code,omitempty"`
	// 分页限制
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 分组ID
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 楼层ID
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 类别ID
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 楼宇ID
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
}

var poolSpaceGroupQuery = sync.Pool{
	New: func() any {
		return new(SpaceGroupQuery)
	},
}

// GetSpaceGroupQuery() 从对象池中获取SpaceGroupQuery
func GetSpaceGroupQuery() *SpaceGroupQuery {
	return poolSpaceGroupQuery.Get().(*SpaceGroupQuery)
}

// ReleaseSpaceGroupQuery 释放SpaceGroupQuery
func ReleaseSpaceGroupQuery(v *SpaceGroupQuery) {
	v.Ids = v.Ids[:0]
	v.TypeCode = ""
	v.Code = ""
	v.Name = ""
	v.NameOrCode = ""
	v.Limit = 0
	v.GroupId = 0
	v.FloorId = 0
	v.CurrentPage = 0
	v.CompanyId = 0
	v.TypeId = 0
	v.CampusId = 0
	v.BuildingId = 0
	poolSpaceGroupQuery.Put(v)
}
