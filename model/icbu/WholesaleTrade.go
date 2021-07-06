package icbu

// WholesaleTrade 结构体
type WholesaleTrade struct {
	// 发货周期，发货时间相关建议使用此项
	DeliverPeriods []DeliverPeriod `json:"deliver_periods,omitempty" xml:"deliver_periods>deliver_period,omitempty"`
	// 最小计量单位，枚举值
	UnitType string `json:"unit_type,omitempty" xml:"unit_type,omitempty"`
	// 销售方式，按件卖(normal)或者按批卖(batch)
	SaleType string `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 价格，单位是美元，精确到小数点后两位，范围是0.01-9999999.00
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 尺寸，单位是厘米，长宽高范围是1-9999999
	PackageSize string `json:"package_size,omitempty" xml:"package_size,omitempty"`
	// 重量，单位是kg，精确到小数点后三位，范围是0.01-9999999.000
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 每批数量，当sale_type=batch时生效，范围是1-99999
	BatchNumber int64 `json:"batch_number,omitempty" xml:"batch_number,omitempty"`
	// 最小起订量，范围是1-99999
	MinOrderQuantity int64 `json:"min_order_quantity,omitempty" xml:"min_order_quantity,omitempty"`
	// 运费模板ID
	ShippingLineTemplateId int64 `json:"shipping_line_template_id,omitempty" xml:"shipping_line_template_id,omitempty"`
	// 备货期，单位是天，范围是1-60
	HandlingTime int64 `json:"handling_time,omitempty" xml:"handling_time,omitempty"`
	// 体积，单位是立方厘米，范围是1-9999999
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
}
