package waybill

// CustomAreaSingleResult 结构体
type CustomAreaSingleResult struct {
	// keys
	Keys []KeyResult `json:"keys,omitempty" xml:"keys>key_result,omitempty"`
	// 自定义区内容的URL
	CustomAreaUrl string `json:"custom_area_url,omitempty" xml:"custom_area_url,omitempty"`
	// 自定义区id
	CustomAreaId int64 `json:"custom_area_id,omitempty" xml:"custom_area_id,omitempty"`
}
