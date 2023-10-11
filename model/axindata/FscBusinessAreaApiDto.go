package axindata

// FscBusinessAreaApiDto 结构体
type FscBusinessAreaApiDto struct {
	// 业务区域ID
	BusinessAreaId string `json:"business_area_id,omitempty" xml:"business_area_id,omitempty"`
	// 业务区域名称
	BusinessAreaName string `json:"business_area_name,omitempty" xml:"business_area_name,omitempty"`
	// 区域类型 1:出境 2:国内
	BusinessAreaRange string `json:"business_area_range,omitempty" xml:"business_area_range,omitempty"`
}
