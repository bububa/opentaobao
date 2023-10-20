package campus

import (
	"sync"
)

// ApiReturnningWrap 结构体
type ApiReturnningWrap struct {
	// version
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// lastUpdateTime
	LastUpdateTime string `json:"last_update_time,omitempty" xml:"last_update_time,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// spaceType
	SpaceType string `json:"space_type,omitempty" xml:"space_type,omitempty"`
	// typeName
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// typeCode
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// 地图图元id
	FtId string `json:"ft_id,omitempty" xml:"ft_id,omitempty"`
	// typeId
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 地图楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
}

var poolApiReturnningWrap = sync.Pool{
	New: func() any {
		return new(ApiReturnningWrap)
	},
}

// GetApiReturnningWrap() 从对象池中获取ApiReturnningWrap
func GetApiReturnningWrap() *ApiReturnningWrap {
	return poolApiReturnningWrap.Get().(*ApiReturnningWrap)
}

// ReleaseApiReturnningWrap 释放ApiReturnningWrap
func ReleaseApiReturnningWrap(v *ApiReturnningWrap) {
	v.Version = ""
	v.Uuid = ""
	v.LastUpdateTime = ""
	v.Name = ""
	v.Code = ""
	v.SpaceType = ""
	v.TypeName = ""
	v.TypeCode = ""
	v.FtId = ""
	v.TypeId = 0
	v.Id = 0
	v.GeoFloorId = 0
	poolApiReturnningWrap.Put(v)
}
