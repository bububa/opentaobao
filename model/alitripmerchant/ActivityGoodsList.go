package alitripmerchant

// ActivityGoodsList 结构体
type ActivityGoodsList struct {
	// 奖品图片
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奖品id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
