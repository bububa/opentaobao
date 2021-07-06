package deliveryvoucher

// RollbackVoucherRequest 结构体
type RollbackVoucherRequest struct {
	// 券信息,最多100条券记录
	VoucherInfos []DeliveryVoucherInfoDto `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
	// 操作时间
	OperateDate string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
	// 扩展参数
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 流水号:唯一，幂等性判断
	OpId string `json:"op_id,omitempty" xml:"op_id,omitempty"`
	// 第三方服务商标识：appkey
	Provider string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 主订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
