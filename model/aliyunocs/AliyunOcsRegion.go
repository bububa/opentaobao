package aliyunocs

// AliyunOcsRegion 结构体
type AliyunOcsRegion struct {
	// cn-hangzhou
	RegionId string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// "B,D"
	ZoneIds string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
	// 杭州
	LocalName string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
}
