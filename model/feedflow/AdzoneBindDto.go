package feedflow

// AdzoneBindDto 结构体
type AdzoneBindDto struct {
	// 资源位id
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 溢价
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
}
