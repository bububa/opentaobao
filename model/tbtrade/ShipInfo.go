package tbtrade

// ShipInfo 结构体
type ShipInfo struct {
	// 发货方式（小写）
	TransportType string `json:"transport_type,omitempty" xml:"transport_type,omitempty"`
}
