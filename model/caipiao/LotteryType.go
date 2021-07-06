package caipiao

// LotteryType 结构体
type LotteryType struct {
	// 彩种名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 彩种ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
