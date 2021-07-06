package retail

// VendingBizOrderDto 结构体
type VendingBizOrderDto struct {
	// 创单时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 设备编码
	DeviceSn string `json:"device_sn,omitempty" xml:"device_sn,omitempty"`
	// 设备名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 设备地址
	DeviceAddress string `json:"device_address,omitempty" xml:"device_address,omitempty"`
	// 外部设备ID
	DeviceUuid string `json:"device_uuid,omitempty" xml:"device_uuid,omitempty"`
	// 设备Code
	DeviceCode string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// 点位名称
	PointName string `json:"point_name,omitempty" xml:"point_name,omitempty"`
	// 外部订单ID
	MainOuterOrderId string `json:"main_outer_order_id,omitempty" xml:"main_outer_order_id,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 支付状态
	StatusName string `json:"status_name,omitempty" xml:"status_name,omitempty"`
	// 支付类型
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 支付类型名称
	PayTypeName string `json:"pay_type_name,omitempty" xml:"pay_type_name,omitempty"`
	// 条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 货道编号
	RoadId string `json:"road_id,omitempty" xml:"road_id,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 金额，单位分
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 淘系订单id
	TbMainOrderId int64 `json:"tb_main_order_id,omitempty" xml:"tb_main_order_id,omitempty"`
	// -20 已退款，-10 交易关闭 ，10 创单 20 已支付  30 已出货  40 交易完成
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
