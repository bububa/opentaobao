package btrip

// OpenIsvBillSettlementBtripTrainRs 结构体
type OpenIsvBillSettlementBtripTrainRs struct {
	// 出行人名称
	TravelerName string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// 退票手续费
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 发车站
	DeptStation string `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	// 到达日期
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 车次类型
	TrainType string `json:"train_type,omitempty" xml:"train_type,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 出发时间
	DeptTime string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	// 部门id
	DepartmentId string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 车次
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 改签手续费
	ChangeFee string `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 超标审批单号
	OverApplyId string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// 审批单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 坐席
	SeatType string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// 票面票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 预订人use id
	BookerId string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// 项目编码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 级联部门
	CascadeDepartment string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	// 订单金额
	OrderPrice string `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// 运行时长
	RunTime string `json:"run_time,omitempty" xml:"run_time,omitempty"`
	// 成本中心编号
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 坐席
	SeatNo string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// 末级部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 服务费,仅在feeType 6007、6008中展示，枚举详见语雀
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 结算类型，枚举详见语雀
	SettlementType string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	// 票价
	TicketPrice string `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 优惠券
	Coupon string `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 到达站
	ArrStation string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// 成本中心名称
	CostCenter string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 支付流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 结算金额
	SettlementFee string `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// 序号
	Index string `json:"index,omitempty" xml:"index,omitempty"`
	// 预定时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 预订人名称
	BookerName string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// 费用类型，枚举详见语雀
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 出行人use id
	TravelerId string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// 资金方向，枚举详见语雀
	CapitalDirection string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	// 结算时间
	SettlementTime string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 出发日期
	DeptDate string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	// 预订人工号
	BookerJobNo string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	// 出行人工号
	TravelerJobNo string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	// 入账时间
	BillRecordTime string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	// 预存赠送金额消费
	SettlementGrantFee string `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 车厢号
	CoachNo string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	// 短票号
	ShortTicketNo string `json:"short_ticket_no,omitempty" xml:"short_ticket_no,omitempty"`
	// 主键id，遇到相同id，已最新为准（数据会更新）
	PrimaryId int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// 入账状态，枚举详见语雀
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 票据类型，枚举详见语雀
	VoucherType int64 `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}
