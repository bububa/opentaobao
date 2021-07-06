package alicom

// DistributionOrderLogisticsModel 结构体
type DistributionOrderLogisticsModel struct {
	// 快递公司编码
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 快递名称
	ExpressName string `json:"express_name,omitempty" xml:"express_name,omitempty"`
	// 快递单号
	ExpressNumber string `json:"express_number,omitempty" xml:"express_number,omitempty"`
	// 产品编码，如ICCID
	ProductSerialNo string `json:"product_serial_no,omitempty" xml:"product_serial_no,omitempty"`
	// 淘宝交易订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 需要变更的订购状态, NOT_ORDER(1, "未订购"), 未订购； ORDER_AUDIT(2, "订购中"), 无订购接口，提交给供应商，线下受理中； ON_ORDER(3, "订购中"), 有订购接口，线上受理中； FAILURE(4, "订购失败")，订购失败； SUCCESS(5, "订购成功")，订购成功； CANCEL(6, "订购取消")，订购取消
	ContractOrderStatus string `json:"contract_order_status,omitempty" xml:"contract_order_status,omitempty"`
	// iccid
	Iccid string `json:"iccid,omitempty" xml:"iccid,omitempty"`
	// 商品编码
	ItemSerialNo string `json:"item_serial_no,omitempty" xml:"item_serial_no,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 失败原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 分销商编号
	DistributorId int64 `json:"distributor_id,omitempty" xml:"distributor_id,omitempty"`
	// 身份证相关信息
	IdCardModel *IdCardModel `json:"id_card_model,omitempty" xml:"id_card_model,omitempty"`
}
