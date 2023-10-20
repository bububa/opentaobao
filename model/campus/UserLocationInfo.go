package campus

import (
	"sync"
)

// UserLocationInfo 结构体
type UserLocationInfo struct {
	// 用户身份信息,mac地址或者支付宝id或者其他
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 结构
	Geometry string `json:"geometry,omitempty" xml:"geometry,omitempty"`
	// 空间单元id
	SpaceUnitId string `json:"space_unit_id,omitempty" xml:"space_unit_id,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 进度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 描述identity的种类,1代表传入的是userid,此种情况identity为空,2代表mac地址,3代表支付宝id
	IdentityType int64 `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// 来源
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 高德楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 坐标系EPSG code
	SRID int64 `json:"s_r_i_d,omitempty" xml:"s_r_i_d,omitempty"`
	// timestamp
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 经纬度类型1代表GPS,2代表mercator;
	GeometryType int64 `json:"geometry_type,omitempty" xml:"geometry_type,omitempty"`
	// 是否删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}

var poolUserLocationInfo = sync.Pool{
	New: func() any {
		return new(UserLocationInfo)
	},
}

// GetUserLocationInfo() 从对象池中获取UserLocationInfo
func GetUserLocationInfo() *UserLocationInfo {
	return poolUserLocationInfo.Get().(*UserLocationInfo)
}

// ReleaseUserLocationInfo 释放UserLocationInfo
func ReleaseUserLocationInfo(v *UserLocationInfo) {
	v.Identity = ""
	v.Geometry = ""
	v.SpaceUnitId = ""
	v.Lat = ""
	v.Lng = ""
	v.UserId = ""
	v.Id = 0
	v.IdentityType = 0
	v.Source = 0
	v.GeoFloorId = 0
	v.CampusId = 0
	v.SRID = 0
	v.Timestamp = 0
	v.GeometryType = 0
	v.IsDelete = false
	poolUserLocationInfo.Put(v)
}
