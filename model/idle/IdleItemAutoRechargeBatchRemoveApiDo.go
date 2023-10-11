package idle

// IdleItemAutoRechargeBatchRemoveApiDo 结构体
type IdleItemAutoRechargeBatchRemoveApiDo struct {
	// 商品id列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
}
