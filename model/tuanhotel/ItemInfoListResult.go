package tuanhotel

// ItemInfoListResult 结构体
type ItemInfoListResult struct {
	// 宝贝信息
	TuanItemOnlines []TuanItemOnlineManagerVo `json:"tuan_item_onlines,omitempty" xml:"tuan_item_onlines>tuan_item_online_manager_vo,omitempty"`
}
