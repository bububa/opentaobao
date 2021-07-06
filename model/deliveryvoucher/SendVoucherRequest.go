package deliveryvoucher

// SendVoucherRequest 结构体
type SendVoucherRequest struct {
	// 券信息,券信息,最多100条券记录
	VoucherInfos []DeliveryVoucherInfoDto `json:"voucher_infos,omitempty" xml:"voucher_infos>delivery_voucher_info_dto,omitempty"`
	// 操作时间
	OperateDate string `json:"operate_date,omitempty" xml:"operate_date,omitempty"`
	// 扩展参数
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 流水号:唯一，幂等性判断
	OpId string `json:"op_id,omitempty" xml:"op_id,omitempty"`
	// 第三方服务商标识：top appkey
	Provider string `json:"provider,omitempty" xml:"provider,omitempty"`
	// 收件人电话
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收件人
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 收件人地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 物流公司编码
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 物流名称
	LogisticsName string `json:"logistics_name,omitempty" xml:"logistics_name,omitempty"`
	// 主订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
