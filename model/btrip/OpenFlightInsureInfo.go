package btrip

// OpenFlightInsureInfo 结构体
type OpenFlightInsureInfo struct {
	// 乘机人(保险人)姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 保单号
	InsureNo string `json:"insure_no,omitempty" xml:"insure_no,omitempty"`
	// 状态：1已出保 2已退保&#39;
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
