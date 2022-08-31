package btrip

// TripPreferenceRq 结构体
type TripPreferenceRq struct {
	// 仓位代码
	CabinList []string `json:"cabin_list,omitempty" xml:"cabin_list>string,omitempty"`
}
