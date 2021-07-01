package hotelalliance

// AllianceInfoResult 结构体
type AllianceInfoResult struct {
	// 菲住hid列表
	AllianceHids []int64 `json:"alliance_hids,omitempty" xml:"alliance_hids>int64,omitempty"`
}
