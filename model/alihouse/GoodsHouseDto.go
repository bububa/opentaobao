package alihouse

// GoodsHouseDto 结构体
type GoodsHouseDto struct {
	// 委托房源所属小区外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 房源外部id
	OuterHouseId string `json:"outer_house_id,omitempty" xml:"outer_house_id,omitempty"`
}
