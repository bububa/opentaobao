package alihouse

// SyncHouseResultDto 结构体
type SyncHouseResultDto struct {
	// 委托版本
	EntrustVersion string `json:"entrust_version,omitempty" xml:"entrust_version,omitempty"`
	// 房源id
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 委托id 公域房源同步时此返回值为空
	HouseId int64 `json:"house_id,omitempty" xml:"house_id,omitempty"`
	// 委托商品
	EntrustItemId int64 `json:"entrust_item_id,omitempty" xml:"entrust_item_id,omitempty"`
	// 小区商品
	CommunityItemId int64 `json:"community_item_id,omitempty" xml:"community_item_id,omitempty"`
	// 小区ID
	CommunityId int64 `json:"community_id,omitempty" xml:"community_id,omitempty"`
}
