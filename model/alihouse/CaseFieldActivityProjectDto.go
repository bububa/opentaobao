package alihouse

// CaseFieldActivityProjectDto 结构体
type CaseFieldActivityProjectDto struct {
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 活动地址
	ActivityAddress string `json:"activity_address,omitempty" xml:"activity_address,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 报名人数上限
	SignUpLimit int64 `json:"sign_up_limit,omitempty" xml:"sign_up_limit,omitempty"`
}
