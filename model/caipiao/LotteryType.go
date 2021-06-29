package caipiao

// LotteryType 
type LotteryType struct {
    // 彩种ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 彩种名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
