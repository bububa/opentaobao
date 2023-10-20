package btrip

import (
	"sync"
)

// OpenApiZzdTrainOrderRs 结构体
type OpenApiZzdTrainOrderRs struct {
	// 成本中心信息
	CostCenterList []CostCenterDo `json:"cost_center_list,omitempty" xml:"cost_center_list>cost_center_do,omitempty"`
	// 用户信息
	ClientInfoDos []ClientInfoDo `json:"client_info_dos,omitempty" xml:"client_info_dos>client_info_do,omitempty"`
	// 票信息，老火车票的票信息为空
	BtripTrainOpenTicketDoList []BtripTrainOpenTicketDo `json:"btrip_train_open_ticket_do_list,omitempty" xml:"btrip_train_open_ticket_do_list>btrip_train_open_ticket_do,omitempty"`
	// 订单状态改变时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单状态：-1, &#34;订单初始化,此状态订单对用户不可见&#34; 0, &#34;待支付&#34; 1, &#34;出票中&#34; 2, &#34;已关闭&#34; 3, &#34;有改签单&#34; 4, &#34;有退票单&#34; 5, &#34;出票完成&#34; 6, &#34;退票申请中&#34; 7, &#34;改签申请中&#34; 9, &#34;出票失败&#34; 10, &#34;改签失败&#34; 11, &#34;退票失败&#34;;
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 申请单标题
	ApplyTitle string `json:"apply_title,omitempty" xml:"apply_title,omitempty"`
	// 第三方申请单id
	ThirdPartApplyId string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 用户名称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第三方企业ID
	ThirdpartCorpId string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// 商旅企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方交易ID
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
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
	// tmc收取的服务费（分）
	ServiceFee int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// tmc收取的退订服务费（分）
	RefundServiceFee int64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	// 退票费（分），退票退回的钱
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 个人支付（分），当pay_type是个人支付时有值
	PersonalAmount int64 `json:"personal_amount,omitempty" xml:"personal_amount,omitempty"`
	// 企业支付（分），当pay_type是企业支付时有值
	CorpAmount int64 `json:"corp_amount,omitempty" xml:"corp_amount,omitempty"`
	// 总结算金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 结算方式:1个人支付，2企业现付，3预存，4月结
	SettleType int64 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// 支付类型,0:企业支付，1:个人支付
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 商旅审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 资金流向:1:支出，2:收入
	TradeAction int64 `json:"trade_action,omitempty" xml:"trade_action,omitempty"`
	// 费用类型 1预定成功2退票成功3改签成功4差额退款6线下退票改签退款7火车票服务费8火车票赔付9火车票改签服务费10出票失败退款
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 差标（分）
	TravelStandardAmount int64 `json:"travel_standard_amount,omitempty" xml:"travel_standard_amount,omitempty"`
	// 改签费（分），改签补的差价
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// tmc收取的改签服务费（分）
	ChangeServiceFee int64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 乘客人数
	PassengerCount int64 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	// 12306收取的改签手续费（分）
	ChangeChargeFee int64 `json:"change_charge_fee,omitempty" xml:"change_charge_fee,omitempty"`
	// 12306收取的退票手续费（分）
	RefundChargeFee int64 `json:"refund_charge_fee,omitempty" xml:"refund_charge_fee,omitempty"`
	// 是否是改签流水
	IsChanged bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
	// 是否是退票流水
	IsRefund bool `json:"is_refund,omitempty" xml:"is_refund,omitempty"`
}

var poolOpenApiZzdTrainOrderRs = sync.Pool{
	New: func() any {
		return new(OpenApiZzdTrainOrderRs)
	},
}

// GetOpenApiZzdTrainOrderRs() 从对象池中获取OpenApiZzdTrainOrderRs
func GetOpenApiZzdTrainOrderRs() *OpenApiZzdTrainOrderRs {
	return poolOpenApiZzdTrainOrderRs.Get().(*OpenApiZzdTrainOrderRs)
}

// ReleaseOpenApiZzdTrainOrderRs 释放OpenApiZzdTrainOrderRs
func ReleaseOpenApiZzdTrainOrderRs(v *OpenApiZzdTrainOrderRs) {
	v.CostCenterList = v.CostCenterList[:0]
	v.ClientInfoDos = v.ClientInfoDos[:0]
	v.BtripTrainOpenTicketDoList = v.BtripTrainOpenTicketDoList[:0]
	v.GmtModified = ""
	v.GmtCreate = ""
	v.Status = ""
	v.ApplyTitle = ""
	v.ThirdPartApplyId = ""
	v.DepartName = ""
	v.UserNick = ""
	v.UserId = ""
	v.ThirdpartCorpId = ""
	v.CorpId = ""
	v.TradeId = ""
	v.Currency = ""
	v.SettlementBatchNo = ""
	v.InvoiceId = ""
	v.InvoiceTitle = ""
	v.ProjectCode = ""
	v.ProjectTitle = ""
	v.ThirdpartItineraryId = ""
	v.ItineraryNo = ""
	v.ServiceFee = 0
	v.RefundServiceFee = 0
	v.RefundFee = 0
	v.PersonalAmount = 0
	v.CorpAmount = 0
	v.Amount = 0
	v.SettleType = 0
	v.PayType = 0
	v.ApplyId = 0
	v.OrderId = 0
	v.TradeAction = 0
	v.TradeType = 0
	v.TravelStandardAmount = 0
	v.ChangeFee = 0
	v.ChangeServiceFee = 0
	v.TicketCount = 0
	v.PassengerCount = 0
	v.ChangeChargeFee = 0
	v.RefundChargeFee = 0
	v.IsChanged = false
	v.IsRefund = false
	poolOpenApiZzdTrainOrderRs.Put(v)
}
