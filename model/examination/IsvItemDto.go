package examination

// IsvItemDto 结构体
type IsvItemDto struct {
	// isv的单项id
	IsvItemId string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
	// 单项描述
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 版本号，防止isv更改未同步给健康，提供给isv做校验的
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 售价，单位分
	SoldPrice string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
}
