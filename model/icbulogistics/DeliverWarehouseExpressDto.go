package icbulogistics

// DeliverWarehouseExpressDto 结构体
type DeliverWarehouseExpressDto struct {
	// 运单号
	TrackingNumbers []string `json:"tracking_numbers,omitempty" xml:"tracking_numbers>string,omitempty"`
	// 国内快递公司code
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 包裹数量
	PackageQuantity string `json:"package_quantity,omitempty" xml:"package_quantity,omitempty"`
}
