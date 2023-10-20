package campus

import (
	"sync"
)

// Content 结构体
type Content struct {
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 类型
	SpaceType string `json:"space_type,omitempty" xml:"space_type,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 修改人
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 楼宇名称
	BuildingName string `json:"building_name,omitempty" xml:"building_name,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 楼层名称
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 高度
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 地图楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
	// 楼宇id
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}

var poolContent = sync.Pool{
	New: func() any {
		return new(Content)
	},
}

// GetContent() 从对象池中获取Content
func GetContent() *Content {
	return poolContent.Get().(*Content)
}

// ReleaseContent 释放Content
func ReleaseContent(v *Content) {
	v.Area = ""
	v.Creator = ""
	v.GmtModified = ""
	v.Code = ""
	v.SpaceType = ""
	v.CampusName = ""
	v.Modifier = ""
	v.Description = ""
	v.GmtCreate = ""
	v.Uuid = ""
	v.BuildingName = ""
	v.Name = ""
	v.FloorName = ""
	v.Height = ""
	v.OrderNo = 0
	v.CampusId = 0
	v.GeoFloorId = 0
	v.BuildingId = 0
	v.FloorId = 0
	v.CompanyId = 0
	v.Id = 0
	v.Status = 0
	v.IsDelete = false
	poolContent.Put(v)
}
