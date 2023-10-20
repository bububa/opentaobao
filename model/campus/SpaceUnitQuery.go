package campus

import (
	"sync"
)

// SpaceUnitQuery 结构体
type SpaceUnitQuery struct {
	// 空间单元ID集合
	Ids []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
	// 分组编码
	GroupCode string `json:"group_code,omitempty" xml:"group_code,omitempty"`
	// 类型编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// 空间单元名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 空间编号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 名称或者编码
	NameOrCode string `json:"name_or_code,omitempty" xml:"name_or_code,omitempty"`
	// 空间编号关键字用于模糊查询
	CodeKeyWord string `json:"code_key_word,omitempty" xml:"code_key_word,omitempty"`
	// 10
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 分组id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 类型id
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 楼层ID
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 楼宇ID
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// 空间单元类别
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 启用停用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSpaceUnitQuery = sync.Pool{
	New: func() any {
		return new(SpaceUnitQuery)
	},
}

// GetSpaceUnitQuery() 从对象池中获取SpaceUnitQuery
func GetSpaceUnitQuery() *SpaceUnitQuery {
	return poolSpaceUnitQuery.Get().(*SpaceUnitQuery)
}

// ReleaseSpaceUnitQuery 释放SpaceUnitQuery
func ReleaseSpaceUnitQuery(v *SpaceUnitQuery) {
	v.Ids = v.Ids[:0]
	v.GroupCode = ""
	v.TypeCode = ""
	v.Name = ""
	v.Code = ""
	v.NameOrCode = ""
	v.CodeKeyWord = ""
	v.Limit = 0
	v.GroupId = 0
	v.CurrentPage = 0
	v.TypeId = 0
	v.FloorId = 0
	v.CompanyId = 0
	v.CampusId = 0
	v.BuildingId = 0
	v.Category = 0
	v.Status = 0
	poolSpaceUnitQuery.Put(v)
}
