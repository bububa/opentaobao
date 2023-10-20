package btrip

import (
	"sync"
)

// OpenApiZzdVehicleOrderRs 结构体
type OpenApiZzdVehicleOrderRs struct {
	// 成本中心信息
	CostCenterList []CostCenterDo `json:"cost_center_list,omitempty" xml:"cost_center_list>cost_center_do,omitempty"`
	// 用户信息
	ClientInfoDos []ClientInfoDo `json:"client_info_dos,omitempty" xml:"client_info_dos>client_info_do,omitempty"`
	// 用车信息
	CarInfoDoList []CarInfoDo `json:"car_info_do_list,omitempty" xml:"car_info_do_list>car_info_do,omitempty"`
	// 第三方交易ID
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 商旅企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方企业id
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 第三方用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 企业维护的用户名称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 第三方申请单id
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 申请单标题
	ApplyTitle string `json:"apply_title,omitempty" xml:"apply_title,omitempty"`
	// 4:已取消 402:取消产生费用自动支付已投诉 6:取消产生费用已确认 403:取消产生费用已确认已投诉 404:取消产生费用未支付已投诉 500:行程结束不认可费用 502:行程结束已支付已投诉 700:行程结束车费合理 701:行程结束自动扣款系统检测合理待确认 702:行程结束自动扣款系统检测合理已确认 703:行程结束自动扣款系统检测费用不合理已投诉 704:历史行程结束待确认订单，订正为此 705:客服关单 706:行程结束自动扣款系统检测合理已确认已投诉 707:行程结束自动扣款系统检测合理直接投诉
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户认可/不认可原因
	UserConfirmReason string `json:"user_confirm_reason,omitempty" xml:"user_confirm_reason,omitempty"`
	// 取消理由
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单状态改变时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// TRAVEL: 差旅, TRAFFIC: 市内交通, WORK: 加班, OTHER: 其它
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 货币种类
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 结算批次：该订单结算时间
	SettlementBatchNo string `json:"settlement_batch_no,omitempty" xml:"settlement_batch_no,omitempty"`
	// 商旅发票ID
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 项目编号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 项目名称
	ProjectTitle string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// 第三方行程单号
	ThirdpartItineraryId string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// 商旅行程单号
	ItineraryNo string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	// 取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// v_sp_t_1:用车里程，v_sp_t_2:实际下车点，v_sp_t_3:用车金额，v_sp_t_4:用车次数，v_sp_t_5:跨城订单
	SpecialTypes string `json:"special_types,omitempty" xml:"special_types,omitempty"`
	// 1. 用车支付 2. 服务费 3. 用车取消后收费 101. 用车退款 102. 用车赔付
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 资金流向:1:支出，2:收入
	TradeAction int64 `json:"trade_action,omitempty" xml:"trade_action,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 用户确认 0未确认 1确认 2 有异议 3 系统自动检测不合理
	UserConfirm int64 `json:"user_confirm,omitempty" xml:"user_confirm,omitempty"`
	// 支付类型,0:企业支付 2.混合支付
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 结算方式:1：个人现付，2:企业现付,4:企业月结，8、企业预存
	SettleType int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// 总结算金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 企业支付（分）
	CorpAmount int64 `json:"corp_amount,omitempty" xml:"corp_amount,omitempty"`
	// 个人支付（分），当pay_type是混合支付时有值
	PersonalAmount int64 `json:"personal_amount,omitempty" xml:"personal_amount,omitempty"`
	// 退票费（分）
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 服务费（分）
	ServiceFee int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 是否特殊订单
	IsSpecial bool `json:"is_special,omitempty" xml:"is_special,omitempty"`
}

var poolOpenApiZzdVehicleOrderRs = sync.Pool{
	New: func() any {
		return new(OpenApiZzdVehicleOrderRs)
	},
}

// GetOpenApiZzdVehicleOrderRs() 从对象池中获取OpenApiZzdVehicleOrderRs
func GetOpenApiZzdVehicleOrderRs() *OpenApiZzdVehicleOrderRs {
	return poolOpenApiZzdVehicleOrderRs.Get().(*OpenApiZzdVehicleOrderRs)
}

// ReleaseOpenApiZzdVehicleOrderRs 释放OpenApiZzdVehicleOrderRs
func ReleaseOpenApiZzdVehicleOrderRs(v *OpenApiZzdVehicleOrderRs) {
	v.CostCenterList = v.CostCenterList[:0]
	v.ClientInfoDos = v.ClientInfoDos[:0]
	v.CarInfoDoList = v.CarInfoDoList[:0]
	v.TradeId = ""
	v.CorpId = ""
	v.ThirdpartCorpId = ""
	v.UserId = ""
	v.UserNick = ""
	v.DepartName = ""
	v.ThirdPartApplyId = ""
	v.ApplyTitle = ""
	v.Status = ""
	v.UserConfirmReason = ""
	v.CancelReason = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.OrderType = ""
	v.Currency = ""
	v.SettlementBatchNo = ""
	v.InvoiceId = ""
	v.InvoiceTitle = ""
	v.ProjectCode = ""
	v.ProjectTitle = ""
	v.ThirdpartItineraryId = ""
	v.ItineraryNo = ""
	v.CancelTime = ""
	v.PayTime = ""
	v.SpecialTypes = ""
	v.TradeType = 0
	v.TradeAction = 0
	v.OrderId = 0
	v.ApplyId = 0
	v.UserConfirm = 0
	v.PayType = 0
	v.SettleType = 0
	v.Amount = 0
	v.CorpAmount = 0
	v.PersonalAmount = 0
	v.RefundFee = 0
	v.ServiceFee = 0
	v.IsSpecial = false
	poolOpenApiZzdVehicleOrderRs.Put(v)
}
