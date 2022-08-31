package btrip

// OpenApiZzdFlightOrderRs 结构体
type OpenApiZzdFlightOrderRs struct {
	// 成本中心信息
	CostCenterList []CostCenterDo `json:"cost_center_list,omitempty" xml:"cost_center_list>cost_center_do,omitempty"`
	// 乘机人列表信息
	ClientInfodos []ClientInfoDo `json:"client_infodos,omitempty" xml:"client_infodos>client_info_do,omitempty"`
	// 票据信息
	Ticketdos []OpenTicketDo `json:"ticketdos,omitempty" xml:"ticketdos>open_ticket_do,omitempty"`
	// 第三方交易号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 商旅企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方企业ID
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 第三方用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第三方用户姓名
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 第三方申请单ID
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 审批单名称
	ApplyTitle string `json:"apply_title,omitempty" xml:"apply_title,omitempty"`
	// 订单状态：0, &#34;待支付&#34;;1, &#34;出票中&#34;;2, &#34;已关闭&#34;;3,&#34;有改签单&#34;;4, &#34;有退票单&#34;;5, &#34;出票成功&#34;;6, &#34;退票申请中&#34;;7, &#34;改签申请中&#34;;10, &#34;订单关闭&#34;;
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单状态改变时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 货币类型
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 审批人名称，多个/分开
	Approver string `json:"approver,omitempty" xml:"approver,omitempty"`
	// 结算批次：corpidyyyymm（该订单结算时间）
	SettlementBatchNo string `json:"settlement_batch_no,omitempty" xml:"settlement_batch_no,omitempty"`
	// 商旅发票ID
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 项目编号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 项目名称
	ProjectTitle string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// 第三方行程id
	ThirdpartItineraryId string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// 行程id
	ItineraryNo string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	// &#39;交易类型：1:机票预订，2:机票改签，3:保险费，4:行程单邮寄费，6：机票退票手续费，101：预订退款，102:改签退款，103:保险退款，104:行程单邮寄退款，105：机票赔付，106:机票改签服务费
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 资金流向:1:支出，2:收入
	TradeAction int64 `json:"trade_action,omitempty" xml:"trade_action,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 申请单ID
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 0:企业支付，1:个人支付
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 结算方式：1：个人现付，2:企业现付,4:企业月结，8、企业预存&#39;
	SettleType int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// 总结算金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 企业支付金额（分），当pay_type是企业支付时有值
	CorpAmount int64 `json:"corp_amount,omitempty" xml:"corp_amount,omitempty"`
	// 个人支付金额（分），当pay_type是个人支付时有值
	PersonalAmount int64 `json:"personal_amount,omitempty" xml:"personal_amount,omitempty"`
	// 退票费（分），改签退款/预订退款 退回的钱
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// tmc收取的退票服务费（分）
	RefundServiceFee int64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	// tmc收取的预订服务费（分）
	ServiceFee int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 差标金额（分）
	TravelStandardAmount int64 `json:"travel_standard_amount,omitempty" xml:"travel_standard_amount,omitempty"`
	// 改签费用（分）
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// tmc收取的改签服务费（分）
	ChangeServiceFee int64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
	// 行程单邮费
	DeliveryFee int64 `json:"delivery_fee,omitempty" xml:"delivery_fee,omitempty"`
	// 保险费用（分）
	InsuranceFee int64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	// 票张数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 乘机人数量
	PassengerCount int64 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	// 航司收取的改签手续费
	ChangeChargeFee int64 `json:"change_charge_fee,omitempty" xml:"change_charge_fee,omitempty"`
	// 航司收取的退票手续费
	RefundChargeFee int64 `json:"refund_charge_fee,omitempty" xml:"refund_charge_fee,omitempty"`
	// 是否是改签流水
	IsChanged bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
	// 是否是退票流水
	IsRefund bool `json:"is_refund,omitempty" xml:"is_refund,omitempty"`
}
