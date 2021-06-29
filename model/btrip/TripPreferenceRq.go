package btrip

// TripPreferenceRq 
type TripPreferenceRq struct {
    // 仓位代码
    CabinList   []string `json:"cabin_list,omitempty" xml:"cabin_list>string,omitempty"`
}
