package feedflow

// AdzoneDto 结构体
type AdzoneDto struct {
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 广告位id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
}
