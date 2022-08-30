package btrip

// OpenIsvBillSettlementFlightRs 结构体
type OpenIsvBillSettlementFlightRs struct {
	// 出行人名称
	TravelerName string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// 舱位
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 基建费
	BuildFee string `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
	// 退款手续费
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 到达日期
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 部门id
	DepartmentId string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 折扣率
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// 改签手续费
	ChangeFee string `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 改签差价
	UpgradeCost string `json:"upgrade_cost,omitempty" xml:"upgrade_cost,omitempty"`
	// 超标审批单号
	OverApplyId string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// 审批单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 预订人use id
	BookerId string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// 低价提醒(与最低价差额)
	MostDifferencePrice string `json:"most_difference_price,omitempty" xml:"most_difference_price,omitempty"`
	// 成本中心编码
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 行程单打印序号
	ItineraryNum string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 服务费，仅在feeType  11001、11002中展示
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 行程单金额
	ItineraryPrice string `json:"itinerary_price,omitempty" xml:"itinerary_price,omitempty"`
	// 到达机场
	ArrStation string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// 成本中心名称
	CostCenter string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 结算金额
	SettlementFee string `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// 序号
	Index string `json:"index,omitempty" xml:"index,omitempty"`
	// 预定时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 费用类型，枚举详见语雀
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 订单金额
	CorpPayOrderFee string `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	// 燃油费
	OilFee string `json:"oil_fee,omitempty" xml:"oil_fee,omitempty"`
	// 商旅价优惠金额
	BtripCouponFee string `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	// 是否重复退
	RepeatRefund string `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	// 起飞城市
	DeptCity string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 到达机场航司三字码
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 行程单号
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// 航司二字码
	AirlineCorpCode string `json:"airline_corp_code,omitempty" xml:"airline_corp_code,omitempty"`
	// 起飞机场
	DeptStation string `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	// 起飞时间
	DeptTime string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	// 低价航班价格
	MostPrice string `json:"most_price,omitempty" xml:"most_price,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 舱位代码
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 起飞机场航司三字码
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 项目编码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 级联部门
	CascadeDepartment string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	// 末级部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 低价提醒（折扣）
	MostDifferenceDiscount string `json:"most_difference_discount,omitempty" xml:"most_difference_discount,omitempty"`
	// 结算类型，枚举详见语雀
	SettlementType string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	// 优惠券
	Coupon string `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 支付流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 航司名称
	AirlineCorpName string `json:"airline_corp_name,omitempty" xml:"airline_corp_name,omitempty"`
	// 低价提醒(航班号)
	MostDifferenceFlightNo string `json:"most_difference_flight_no,omitempty" xml:"most_difference_flight_no,omitempty"`
	// 预订人名称
	BookerName string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// 改签退票手续费
	RefundUpgradeCost string `json:"refund_upgrade_cost,omitempty" xml:"refund_upgrade_cost,omitempty"`
	// 出行人use id
	TravelerId string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// 资金方向，枚举详见语雀
	CapitalDirection string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	// 保险费
	InsuranceFee string `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	// 销售金额
	SealPrice string `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	// 结算时间
	SettlementTime string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// 低价提醒（起飞时间）
	MostDifferenceDeptTime string `json:"most_difference_dept_time,omitempty" xml:"most_difference_dept_time,omitempty"`
	// 不选低价原因
	MostDifferenceReason string `json:"most_difference_reason,omitempty" xml:"most_difference_reason,omitempty"`
	// 协议价优惠金额
	NegotiationCouponFee string `json:"negotiation_coupon_fee,omitempty" xml:"negotiation_coupon_fee,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 起飞日期
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
	// 主键id，遇到相同id，已最新为准（数据会更新）
	PrimaryId int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// 入账状态，枚举详见语雀
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 提前预定天数
	AdvanceDay int64 `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	// 票据类型，枚举详见语雀
	VoucherType int64 `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}
