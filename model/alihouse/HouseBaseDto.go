package alihouse

// HouseBaseDto 结构体
type HouseBaseDto struct {
	// 图片信息
	Pictures []ProjectPictureDto `json:"pictures,omitempty" xml:"pictures>project_picture_dto,omitempty"`
	// 外部小区id
	OuterCommunityId string `json:"outer_community_id,omitempty" xml:"outer_community_id,omitempty"`
	// 外部房屋id
	OuterHouseBaseId string `json:"outer_house_base_id,omitempty" xml:"outer_house_base_id,omitempty"`
	// 外部公司id
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// e码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 楼栋号
	BuildingNo string `json:"building_no,omitempty" xml:"building_no,omitempty"`
	// 户室号
	UnitNo string `json:"unit_no,omitempty" xml:"unit_no,omitempty"`
	// 房间号
	RoomNo string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// 室内面积
	InsideArea string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
	// 建筑面积
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 房屋室数
	Room int64 `json:"room,omitempty" xml:"room,omitempty"`
	// 房屋厅数
	Hall int64 `json:"hall,omitempty" xml:"hall,omitempty"`
	// 卫生间数
	Toilet int64 `json:"toilet,omitempty" xml:"toilet,omitempty"`
	// 厨房数
	Kitchen int64 `json:"kitchen,omitempty" xml:"kitchen,omitempty"`
	// 阳台数
	Balcony int64 `json:"balcony,omitempty" xml:"balcony,omitempty"`
	// 所在楼层
	Floor int64 `json:"floor,omitempty" xml:"floor,omitempty"`
	// 总楼层数
	TotalFloor int64 `json:"total_floor,omitempty" xml:"total_floor,omitempty"`
	// 装修情况
	Fitment int64 `json:"fitment,omitempty" xml:"fitment,omitempty"`
	// 供暖类型（0-未知 1-自采暖 2-集中供暖 3-无供暖）
	HeatingType int64 `json:"heating_type,omitempty" xml:"heating_type,omitempty"`
	// 朝向
	Direction int64 `json:"direction,omitempty" xml:"direction,omitempty"`
	// 户型结构
	HouseTypeStructure int64 `json:"house_type_structure,omitempty" xml:"house_type_structure,omitempty"`
	// 车位
	ParkingType int64 `json:"parking_type,omitempty" xml:"parking_type,omitempty"`
	// 用水
	WaterType int64 `json:"water_type,omitempty" xml:"water_type,omitempty"`
	// 用电
	PowerType int64 `json:"power_type,omitempty" xml:"power_type,omitempty"`
	// 燃气
	GasType int64 `json:"gas_type,omitempty" xml:"gas_type,omitempty"`
}
