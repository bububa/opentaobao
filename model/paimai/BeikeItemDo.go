package paimai

// BeikeItemDo 结构体
type BeikeItemDo struct {
	// 城市（简称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 房屋用途
	HousePurpose string `json:"house_purpose,omitempty" xml:"house_purpose,omitempty"`
	// 交易权属
	TransOwnership string `json:"trans_ownership,omitempty" xml:"trans_ownership,omitempty"`
	// 特性
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 物业公司
	PropertyManagement string `json:"property_management,omitempty" xml:"property_management,omitempty"`
	// 小区别名
	CommunityAlias string `json:"community_alias,omitempty" xml:"community_alias,omitempty"`
	// 经纪人电话
	BrokerTele string `json:"broker_tele,omitempty" xml:"broker_tele,omitempty"`
	// 朝向
	HouseToward string `json:"house_toward,omitempty" xml:"house_toward,omitempty"`
	// 楼层
	FloorDesc string `json:"floor_desc,omitempty" xml:"floor_desc,omitempty"`
	// 省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 车位比
	ParkingRatio string `json:"parking_ratio,omitempty" xml:"parking_ratio,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 户型
	HouseType string `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 小区
	Community string `json:"community,omitempty" xml:"community,omitempty"`
	// 容积率
	VolumeRate string `json:"volume_rate,omitempty" xml:"volume_rate,omitempty"`
	// 主图url
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 商圈
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 区位码
	Location string `json:"location,omitempty" xml:"location,omitempty"`
	// 建筑类型
	BuildingType string `json:"building_type,omitempty" xml:"building_type,omitempty"`
	// 小区参考均价（单位：元）Double
	RefUnitPrice int64 `json:"ref_unit_price,omitempty" xml:"ref_unit_price,omitempty"`
	// 总价（单位：元）Double
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 小区评分
	CommunityScore int64 `json:"community_score,omitempty" xml:"community_score,omitempty"`
	// 建成年代
	BuildingTime int64 `json:"building_time,omitempty" xml:"building_time,omitempty"`
	// 面积
	HouseArea int64 `json:"house_area,omitempty" xml:"house_area,omitempty"`
	// 物业费
	PropertyUnitPrice int64 `json:"property_unit_price,omitempty" xml:"property_unit_price,omitempty"`
	// 经度
	Lat int64 `json:"lat,omitempty" xml:"lat,omitempty"`
	// 单价
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 纬度
	Lng int64 `json:"lng,omitempty" xml:"lng,omitempty"`
	// 在售房屋套数
	HousesForSale int64 `json:"houses_for_sale,omitempty" xml:"houses_for_sale,omitempty"`
	// id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 总楼层
	TotalFloor int64 `json:"total_floor,omitempty" xml:"total_floor,omitempty"`
	// 在租房屋套数
	HousesForRent int64 `json:"houses_for_rent,omitempty" xml:"houses_for_rent,omitempty"`
	// 绿化率
	GreeningRate int64 `json:"greening_rate,omitempty" xml:"greening_rate,omitempty"`
}
