package alihouse

// QueryStorePunishDto 结构体
type QueryStorePunishDto struct {
	// 外部门店id列表
	OuterStoreIds []string `json:"outer_store_ids,omitempty" xml:"outer_store_ids>string,omitempty"`
}
