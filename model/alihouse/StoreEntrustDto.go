package alihouse

// StoreEntrustDto 结构体
type StoreEntrustDto struct {
	// 房源委托版本
	EntrustVersion string `json:"entrust_version,omitempty" xml:"entrust_version,omitempty"`
	// 房源套内面积
	InsideArea string `json:"inside_area,omitempty" xml:"inside_area,omitempty"`
	// 房源建筑面积
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 房号
	RoomNo string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// 单元号
	UnitNo string `json:"unit_no,omitempty" xml:"unit_no,omitempty"`
	// 楼栋号
	BuildingNo string `json:"building_no,omitempty" xml:"building_no,omitempty"`
	// 房源详细地址
	HouseAddress string `json:"house_address,omitempty" xml:"house_address,omitempty"`
	// 小区名称
	CommunityName string `json:"community_name,omitempty" xml:"community_name,omitempty"`
	// 房源商品id
	EntrustItemId int64 `json:"entrust_item_id,omitempty" xml:"entrust_item_id,omitempty"`
	// 委托id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 小区商品id
	CommunityItemId int64 `json:"community_item_id,omitempty" xml:"community_item_id,omitempty"`
	// 小区id
	CommunityId int64 `json:"community_id,omitempty" xml:"community_id,omitempty"`
	// 交易状态
	TradeStatus int64 `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
	// 是否存在房源商品
	IsExist bool `json:"is_exist,omitempty" xml:"is_exist,omitempty"`
	// 是否存在委托
	IsCommission bool `json:"is_commission,omitempty" xml:"is_commission,omitempty"`
}
