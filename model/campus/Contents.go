package campus

// Contents 结构体
type Contents struct {
	// 同businessId
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建者
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改时间
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 楼宇名称
	BuildingName string `json:"building_name,omitempty" xml:"building_name,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 分类编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 楼层名称
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 分类名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 园区编码
	CampusCode string `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 空间单元名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地图图元id
	FtId string `json:"ft_id,omitempty" xml:"ft_id,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// heightStr
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// areaStr
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 1室内，2室外，3 逻辑, 4 虚拟
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 状态,0为停用,1为启用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 空间种类，1室内，2室外，3 逻辑
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 楼宇id
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// 针对同一位置的同业务poi设置group_id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 分类id
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 地图楼层id
	GeoFloorId int64 `json:"geo_floor_id,omitempty" xml:"geo_floor_id,omitempty"`
	// 是否删除,0代表未删除,1代表删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}
