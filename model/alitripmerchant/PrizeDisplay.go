package alitripmerchant

// PrizeDisplay 结构体
type PrizeDisplay struct {
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奖品URL
	Image string `json:"image,omitempty" xml:"image,omitempty"`
}
