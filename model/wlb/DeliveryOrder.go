package wlb

// DeliveryOrder 结构体
type DeliveryOrder struct {
	// 订单信息
	OrderLine []OrderLine `json:"order_line,omitempty" xml:"order_line>order_line,omitempty"`
	// 订单行
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 发货单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// ERP出库单号,ERP系统内本次出库的唯一标示
	DeliveryOrderCode string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
	// 物流公司编码
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 交接入库单号,例如唯品会入库单号或者门店收货单号、商家仓入库单号等
	RelInBoundOrderCode string `json:"rel_in_bound_order_code,omitempty" xml:"rel_in_bound_order_code,omitempty"`
	// 发货仓库
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 单据类型,出库单类型(JYCK=一般交易出库单;HHCK=换货出库单;BFCK=补发出库单;QTCK=其他出库单;TOBCK=TOB出库;BIGTOBCK=大B2B发货)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 到货渠道类型，VIP＝1、门店＝2、经销商＝3、大润发=4、猫超=5、零售通=6、AE=7、京东=8、亚马逊=9、经销=10、代理=11、商超=12、其他=99
	ArriveChannelType string `json:"arrive_channel_type,omitempty" xml:"arrive_channel_type,omitempty"`
	// 物流公司名称
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 最晚到货时间
	LastArriveDate string `json:"last_arrive_date,omitempty" xml:"last_arrive_date,omitempty"`
	// 扩展信息
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 收货时间区间
	SignTime string `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 是否自提
	IsSelfLifting string `json:"is_self_lifting,omitempty" xml:"is_self_lifting,omitempty"`
	// 配送方式，1=整车直送、2=拼车直送、3=零担、4=快递、5=自提
	TransportMode string `json:"transport_mode,omitempty" xml:"transport_mode,omitempty"`
	// 物流单号
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// 收货人信息
	ReceiverInfo *ReceiverInfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
}
