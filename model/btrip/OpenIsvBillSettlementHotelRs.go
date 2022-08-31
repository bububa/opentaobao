package btrip

// OpenIsvBillSettlementHotelRs 结构体
type OpenIsvBillSettlementHotelRs struct {
	// 出行人名称
	TravelerName string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// 订单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 杂费
	Fees string `json:"fees,omitempty" xml:"fees,omitempty"`
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 部门id
	DepartmentId string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 超标审批id
	OverApplyId string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// 审批id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 预定人use id
	BookerId string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// 成本中心编码
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 房型
	RoomType string `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// 优惠券
	PromotionFee string `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	// 服务费,仅在 feeType 20111、20112中展示，枚举详见语雀
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 房价
	RoomPrice string `json:"room_price,omitempty" xml:"room_price,omitempty"`
	// 成本中心名称
	CostCenter string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 个人退款金额
	PersonRefundFee string `json:"person_refund_fee,omitempty" xml:"person_refund_fee,omitempty"`
	// 结算金额
	SettlementFee string `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// 序号
	Index string `json:"index,omitempty" xml:"index,omitempty"`
	// 预定时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 费用类型，枚举详见语雀
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 离店时间
	CheckoutDate string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	// 个人结算费用
	PersonSettlePrice string `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	// 是否协议价
	IsNegotiation string `json:"is_negotiation,omitempty" xml:"is_negotiation,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 入住城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 入住城市编码，枚举详见语雀
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 出行人工号
	TravelerJobNo string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 项目编码
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 级联部门
	CascadeDepartment string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	// 订单金额
	OrderPrice string `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// 末级部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 企业支付金额
	CorpTotalFee string `json:"corp_total_fee,omitempty" xml:"corp_total_fee,omitempty"`
	// 结算类型，枚举详见语雀
	SettlementType string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	// 支付流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 企业退款金额
	CorpRefundFee string `json:"corp_refund_fee,omitempty" xml:"corp_refund_fee,omitempty"`
	// 福豆支付
	FuPointFee string `json:"fu_point_fee,omitempty" xml:"fu_point_fee,omitempty"`
	// 预订人名称
	BookerName string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 出行人use id
	TravelerId string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// 资金方向，枚举详见语雀
	CapitalDirection string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	// 结算时间
	SettlementTime string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// 预定人工号
	BookerJobNo string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	// 是否合住
	IsShareStr string `json:"is_share_str,omitempty" xml:"is_share_str,omitempty"`
	// 入账时间
	BillRecordTime string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	// 预存赠送金额消费
	SettlementGrantFee string `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 主键id，遇到相同id，已最新为准（数据会更新）
	PrimaryId int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// 总间夜数
	TotalNights int64 `json:"total_nights,omitempty" xml:"total_nights,omitempty"`
	// 入账状态，枚举详见语雀
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 房间号
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 入住天数
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 票据类型，枚举详见语雀
	VoucherType int64 `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}
