package traveltrade

// TopTripBookInfoResult 结构体
type TopTripBookInfoResult struct {
	// 交易预定详情
	Module *AlitripTravelBookinfoQueryModule `json:"module,omitempty" xml:"module,omitempty"`
}
