package btrip

// TrainTransferInfo 结构体
type TrainTransferInfo struct {
	// 出发站名
	FromStationName string `json:"fromStationName,omitempty" xml:"fromStationName,omitempty"`
	// 到达站名
	ToStationName string `json:"toStationName,omitempty" xml:"toStationName,omitempty"`
	// 出发城市名
	FromCityName string `json:"fromCityName,omitempty" xml:"fromCityName,omitempty"`
	// 到达城市名
	ToCityName string `json:"toCityName,omitempty" xml:"toCityName,omitempty"`
	// 起始时间
	StartTime string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 到达时间
	EndTime string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 历史总时长，分钟
	CostTime string `json:"costTime,omitempty" xml:"costTime,omitempty"`
	// 换乘时长，分钟
	WaitTime string `json:"waitTime,omitempty" xml:"waitTime,omitempty"`
	// 中转城市
	MiddleCity string `json:"middleCity,omitempty" xml:"middleCity,omitempty"`
	// 中转站
	MiddleStation string `json:"middleStation,omitempty" xml:"middleStation,omitempty"`
	// 换乘模式  1：同站，2：异站，3：同车
	MiddleType string `json:"middleType,omitempty" xml:"middleType,omitempty"`
	// 中转日期，yyyy-MM-dd
	MiddleDate string `json:"middleDate,omitempty" xml:"middleDate,omitempty"`
}
