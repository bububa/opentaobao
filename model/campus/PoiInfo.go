package campus

import (
	"sync"
)

// PoiInfo 结构体
type PoiInfo struct {
	// spaceGroups
	SpaceGroups []SpaceGroup `json:"space_groups,omitempty" xml:"space_groups>space_group,omitempty"`
	// attrs
	Attrs []TypeAttrInstance `json:"attrs,omitempty" xml:"attrs>type_attr_instance,omitempty"`
	// ftId
	FtId string `json:"ft_id,omitempty" xml:"ft_id,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 空间单元名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 园区编码
	CampusCode string `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
	// 类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 高度
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 楼层名
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 类型编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// 楼宇名称
	BuildingName string `json:"building_name,omitempty" xml:"building_name,omitempty"`
	// businessId
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 修改者
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建者
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 室内/室外/逻辑/虚拟
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 高德楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// category
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 类型id
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 楼宇id
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 园区id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 中心点
	Central *Point `json:"central,omitempty" xml:"central,omitempty"`
	// groupId
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 是否删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}

var poolPoiInfo = sync.Pool{
	New: func() any {
		return new(PoiInfo)
	},
}

// GetPoiInfo() 从对象池中获取PoiInfo
func GetPoiInfo() *PoiInfo {
	return poolPoiInfo.Get().(*PoiInfo)
}

// ReleasePoiInfo 释放PoiInfo
func ReleasePoiInfo(v *PoiInfo) {
	v.SpaceGroups = v.SpaceGroups[:0]
	v.Attrs = v.Attrs[:0]
	v.FtId = ""
	v.Uuid = ""
	v.Name = ""
	v.CampusName = ""
	v.CampusCode = ""
	v.TypeName = ""
	v.Area = ""
	v.Height = ""
	v.FloorName = ""
	v.TypeCode = ""
	v.BuildingName = ""
	v.Code = ""
	v.GmtModified = ""
	v.Modifier = ""
	v.Creator = ""
	v.GmtCreate = ""
	v.CategoryName = ""
	v.GeoFloorId = 0
	v.FloorId = 0
	v.Category = 0
	v.Id = 0
	v.TypeId = 0
	v.CampusId = 0
	v.BuildingId = 0
	v.Status = 0
	v.CompanyId = 0
	v.Central = nil
	v.GroupId = 0
	v.IsDelete = false
	poolPoiInfo.Put(v)
}
