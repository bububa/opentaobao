package lstpos

// CashierFlowDto 结构体
type CashierFlowDto struct {
	// 收银商品快照对象列表
	CashierGoodsDetailDTOList []CashierGoodsDetailDto `json:"cashier_goods_detail_d_t_o_list,omitempty" xml:"cashier_goods_detail_d_t_o_list>cashier_goods_detail_dto,omitempty"`
	// 支付状态 请求失败:requestFail 支付失败:payFail 支付超时:payTimeOut 支付成功:paySuccess
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 支付方式 支付宝:alipay 微信:wechat 现金:cashPay 其他:other
	PayType string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 结算单总价(实收)，单位:分
	TotalActualPrice int64 `json:"total_actual_price,omitempty" xml:"total_actual_price,omitempty"`
	// 结算单总价(应收)，单位:分
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 最后修改 精确到毫秒
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间 精确到毫秒
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// isv订单流水号
	IsvTradeFlowNum string `json:"isv_trade_flow_num,omitempty" xml:"isv_trade_flow_num,omitempty"`
	// 关联lst原始订单号:主要是退款订单关联原始订单号使用
	OriginalTradeFlowNum string `json:"original_trade_flow_num,omitempty" xml:"original_trade_flow_num,omitempty"`
	// 订单类型 0: 交易订单快照  1:退款订单快照  缺省为0
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 数据来源的设备类型   0：pos，1：自动售货机，9999：其它  缺省值：0
	DeviceType int64 `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 设备物理硬件ID（自身保证唯一性）
	HardwareId string `json:"hardware_id,omitempty" xml:"hardware_id,omitempty"`
	// 收银支付明细列表对象
	CashierPayDetailDTOList []CashierPayDetailDto `json:"cashier_pay_detail_d_t_o_list,omitempty" xml:"cashier_pay_detail_d_t_o_list>cashier_pay_detail_dto,omitempty"`
	// 设备品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 设备型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
