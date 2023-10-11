package tblogistics

// ConfirmPackageOrderRequest 结构体
type ConfirmPackageOrderRequest struct {
	// 包裹信息
	Packages []PackageDto `json:"packages,omitempty" xml:"packages>package_dto,omitempty"`
	// 包裹入库单号
	DeliveryOrderCode string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
	// 业务类型，取值：JYPKZXCK(集运包裹正向出库)、JYPKNXCK(集运包裹逆向出库)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 单据状态，取值：confirm(出库确认)、exception(出库异常)
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 业务操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 物流主体，例如：TaoTian(淘天)
	LogisticsOwner string `json:"logistics_owner,omitempty" xml:"logistics_owner,omitempty"`
	// 异常类型，回传出库异常时必传 【包裹破损】 【撤单拦截】 【出库单异常】 【联系不上商家】 【商家拒绝退回】 【其他原因】
	ErrorType string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 异常原因说明，回传出库异常时必传
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 服务商仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// WMS系统作业单号
	OuterOrderCode string `json:"outer_order_code,omitempty" xml:"outer_order_code,omitempty"`
	// 包裹入库单号
	EntryOrderCode string `json:"entry_order_code,omitempty" xml:"entry_order_code,omitempty"`
}
