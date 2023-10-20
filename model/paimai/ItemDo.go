package paimai

import (
	"sync"
)

// ItemDo 结构体
type ItemDo struct {
	// 小区参考均价（单位：元）Double
	RefUnitPrice string `json:"ref_unit_price,omitempty" xml:"ref_unit_price,omitempty"`
	// 城市（简称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 总价（单位：元）Double
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 房屋用途
	HousePurpose string `json:"house_purpose,omitempty" xml:"house_purpose,omitempty"`
	// 交易权属
	TransOwnership string `json:"trans_ownership,omitempty" xml:"trans_ownership,omitempty"`
	// 小区评分
	CommunityScore string `json:"community_score,omitempty" xml:"community_score,omitempty"`
	// 特性
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 物业公司
	PropertyManagement string `json:"property_management,omitempty" xml:"property_management,omitempty"`
	// 小区别名
	CommunityAlias string `json:"community_alias,omitempty" xml:"community_alias,omitempty"`
	// 经纪人电话
	BrokerTele string `json:"broker_tele,omitempty" xml:"broker_tele,omitempty"`
	// 面积
	HouseArea string `json:"house_area,omitempty" xml:"house_area,omitempty"`
	// 朝向
	HouseToward string `json:"house_toward,omitempty" xml:"house_toward,omitempty"`
	// 楼层
	FloorDesc string `json:"floor_desc,omitempty" xml:"floor_desc,omitempty"`
	// 省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 物业费
	PropertyUnitPrice string `json:"property_unit_price,omitempty" xml:"property_unit_price,omitempty"`
	// 经度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 车位比
	ParkingRatio string `json:"parking_ratio,omitempty" xml:"parking_ratio,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 纬度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
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
	// 绿化率
	GreeningRate string `json:"greening_rate,omitempty" xml:"greening_rate,omitempty"`
	// 建筑类型
	BuildingType string `json:"building_type,omitempty" xml:"building_type,omitempty"`
	// 建成年代
	BuildingTime int64 `json:"building_time,omitempty" xml:"building_time,omitempty"`
	// 在售房屋套数
	HousesForSale int64 `json:"houses_for_sale,omitempty" xml:"houses_for_sale,omitempty"`
	// id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 总楼层
	TotalFloor int64 `json:"total_floor,omitempty" xml:"total_floor,omitempty"`
	// 在租房屋套数
	HousesForRent int64 `json:"houses_for_rent,omitempty" xml:"houses_for_rent,omitempty"`
}

var poolItemDo = sync.Pool{
	New: func() any {
		return new(ItemDo)
	},
}

// GetItemDo() 从对象池中获取ItemDo
func GetItemDo() *ItemDo {
	return poolItemDo.Get().(*ItemDo)
}

// ReleaseItemDo 释放ItemDo
func ReleaseItemDo(v *ItemDo) {
	v.RefUnitPrice = ""
	v.City = ""
	v.TotalPrice = ""
	v.Title = ""
	v.HousePurpose = ""
	v.TransOwnership = ""
	v.CommunityScore = ""
	v.Feature = ""
	v.PropertyManagement = ""
	v.CommunityAlias = ""
	v.BrokerTele = ""
	v.HouseArea = ""
	v.HouseToward = ""
	v.FloorDesc = ""
	v.Prov = ""
	v.PropertyUnitPrice = ""
	v.Lat = ""
	v.ParkingRatio = ""
	v.UnitPrice = ""
	v.Address = ""
	v.Lng = ""
	v.HouseType = ""
	v.Community = ""
	v.VolumeRate = ""
	v.ImgUrl = ""
	v.District = ""
	v.Location = ""
	v.GreeningRate = ""
	v.BuildingType = ""
	v.BuildingTime = 0
	v.HousesForSale = 0
	v.ItemId = 0
	v.TotalFloor = 0
	v.HousesForRent = 0
	poolItemDo.Put(v)
}
