package btrip

import (
	"sync"
)

// OpenIsvBillSettlementCarRs 结构体
type OpenIsvBillSettlementCarRs struct {
	// 出行人名称
	TravelerName string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// 到达日期
	ArrDate string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 部门id
	DepartmentId string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// 用车事由
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 超标审批单号
	OverApplyId string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// 审批单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 预订人use id
	BookerId string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// 成本中心编号
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 供应商车型
	ProviderName string `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	// 服务费,仅在feeType 40111 中展示，枚举详见语雀
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 实际行驶公里数
	RealDriveDistance string `json:"real_drive_distance,omitempty" xml:"real_drive_distance,omitempty"`
	// 成本中心名称
	CostCenter string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 车型
	CarLevel string `json:"car_level,omitempty" xml:"car_level,omitempty"`
	// 特别关注订单
	SpecialOrder string `json:"special_order,omitempty" xml:"special_order,omitempty"`
	// 结算金额
	SettlementFee string `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// 序号
	Index string `json:"index,omitempty" xml:"index,omitempty"`
	// 预定时间
	BookTime string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// 费用类型。枚举详见语雀
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 个人支付金额
	PersonSettleFee string `json:"person_settle_fee,omitempty" xml:"person_settle_fee,omitempty"`
	// 员工是否认可
	UserConfirmDesc string `json:"user_confirm_desc,omitempty" xml:"user_confirm_desc,omitempty"`
	// 特别关注原因
	SpecialReason string `json:"special_reason,omitempty" xml:"special_reason,omitempty"`
	// 出发城市
	DeptCity string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 出发时间
	DeptTime string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出行人工号
	TravelerJobNo string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	// 到达地
	ArrLocation string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// 实际上车点
	RealFromAddr string `json:"real_from_addr,omitempty" xml:"real_from_addr,omitempty"`
	// 项目编号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 级联部门
	CascadeDepartment string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	// 订单金额
	OrderPrice string `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// 末级部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 结算方式，枚举详见语雀
	SettlementType string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	// 优惠券
	Coupon string `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// 支付流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 预订人名称
	BookerName string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// 预估公里数
	EstimateDriveDistance string `json:"estimate_drive_distance,omitempty" xml:"estimate_drive_distance,omitempty"`
	// 出行人use id
	TravelerId string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// 资金方向，枚举详见语雀
	CapitalDirection string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	// 实际下车点
	RealToAddr string `json:"real_to_addr,omitempty" xml:"real_to_addr,omitempty"`
	// 结算时间
	SettlementTime string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// 出发点
	DeptLocation string `json:"dept_location,omitempty" xml:"dept_location,omitempty"`
	// 用车原因
	BusinessCategory string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	// 商旅优惠金额
	CouponPrice string `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	// 预定人工号
	BookerJobNo string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	// 预估金额
	EstimatePrice string `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 出发日期
	DeptDate string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	// 子订单号
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
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
	// 票据类型，枚举详见语雀
	VoucherType int64 `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

var poolOpenIsvBillSettlementCarRs = sync.Pool{
	New: func() any {
		return new(OpenIsvBillSettlementCarRs)
	},
}

// GetOpenIsvBillSettlementCarRs() 从对象池中获取OpenIsvBillSettlementCarRs
func GetOpenIsvBillSettlementCarRs() *OpenIsvBillSettlementCarRs {
	return poolOpenIsvBillSettlementCarRs.Get().(*OpenIsvBillSettlementCarRs)
}

// ReleaseOpenIsvBillSettlementCarRs 释放OpenIsvBillSettlementCarRs
func ReleaseOpenIsvBillSettlementCarRs(v *OpenIsvBillSettlementCarRs) {
	v.TravelerName = ""
	v.ArrDate = ""
	v.OrderId = ""
	v.DepartmentId = ""
	v.Memo = ""
	v.OverApplyId = ""
	v.ApplyId = ""
	v.BookerId = ""
	v.CostCenterNumber = ""
	v.InvoiceTitle = ""
	v.ProviderName = ""
	v.ServiceFee = ""
	v.RealDriveDistance = ""
	v.CostCenter = ""
	v.CarLevel = ""
	v.SpecialOrder = ""
	v.SettlementFee = ""
	v.Index = ""
	v.BookTime = ""
	v.FeeType = ""
	v.PersonSettleFee = ""
	v.UserConfirmDesc = ""
	v.SpecialReason = ""
	v.DeptCity = ""
	v.ProjectName = ""
	v.DeptTime = ""
	v.ArrCity = ""
	v.TravelerJobNo = ""
	v.ArrLocation = ""
	v.RealFromAddr = ""
	v.ProjectCode = ""
	v.CascadeDepartment = ""
	v.OrderPrice = ""
	v.Department = ""
	v.SettlementType = ""
	v.Coupon = ""
	v.AlipayTradeNo = ""
	v.BookerName = ""
	v.EstimateDriveDistance = ""
	v.TravelerId = ""
	v.CapitalDirection = ""
	v.RealToAddr = ""
	v.SettlementTime = ""
	v.DeptLocation = ""
	v.BusinessCategory = ""
	v.CouponPrice = ""
	v.BookerJobNo = ""
	v.EstimatePrice = ""
	v.ArrTime = ""
	v.DeptDate = ""
	v.SubOrderId = ""
	v.BillRecordTime = ""
	v.SettlementGrantFee = ""
	v.Remark = ""
	v.PrimaryId = 0
	v.Status = 0
	v.VoucherType = 0
	poolOpenIsvBillSettlementCarRs.Put(v)
}
