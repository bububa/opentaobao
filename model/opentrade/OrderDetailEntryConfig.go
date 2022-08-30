package opentrade

// OrderDetailEntryConfig 结构体
type OrderDetailEntryConfig struct {
	// 跳转入口文案
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 允许展示入口的开始时间，未配置表示一直显示入口
	ExposeStartTime int64 `json:"expose_start_time,omitempty" xml:"expose_start_time,omitempty"`
	// 允许展示入口的结束时间，未配置表示一直显示入口
	ExposeEndTime int64 `json:"expose_end_time,omitempty" xml:"expose_end_time,omitempty"`
}
