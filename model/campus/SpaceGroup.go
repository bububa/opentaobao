package campus

// SpaceGroup 结构体
type SpaceGroup struct {
	// attrs
	Attrs []TypeAttrInstance `json:"attrs,omitempty" xml:"attrs>type_attr_instance,omitempty"`
	// campusName
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// typeName
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// typeCode
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// bigTypeName
	BigTypeName string `json:"big_type_name,omitempty" xml:"big_type_name,omitempty"`
	// buildingName
	BuildingName string `json:"building_name,omitempty" xml:"building_name,omitempty"`
	// description
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// floorName
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// modifier
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// creator
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 地图图元id
	FtId string `json:"ft_id,omitempty" xml:"ft_id,omitempty"`
	// campusId
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// typeId
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// bigTypeId
	BigTypeId int64 `json:"big_type_id,omitempty" xml:"big_type_id,omitempty"`
	// buildingId
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// floorId
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// typeWrap
	TypeWrap *PoiTypeWrap `json:"type_wrap,omitempty" xml:"type_wrap,omitempty"`
	// 地图楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
	// isDelete
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}
