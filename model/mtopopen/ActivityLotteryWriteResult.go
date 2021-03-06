package mtopopen

// ActivityLotteryWriteResult 结构体
type ActivityLotteryWriteResult struct {
	// isv活动url
	H5Url string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
	// isv活动的id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
