package ieagency

import (
	"sync"
)

// IeChangeOrderVo 结构体
type IeChangeOrderVo struct {
	// 乘机人信息
	ChangePassengerVOS []IeChangePassengerVo `json:"change_passenger_v_o_s,omitempty" xml:"change_passenger_v_o_s>ie_change_passenger_vo,omitempty"`
	// 改签票信息
	ChangeTicketVOS []IeChangeTicketVo `json:"change_ticket_v_o_s,omitempty" xml:"change_ticket_v_o_s>ie_change_ticket_vo,omitempty"`
	// 支付宝交易号
	AlipayTradeNO string `json:"alipay_trade_n_o,omitempty" xml:"alipay_trade_n_o,omitempty"`
	// 申请支付时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 买家改签意向备注
	BuyerIntensionMemo string `json:"buyer_intension_memo,omitempty" xml:"buyer_intension_memo,omitempty"`
	// 验真通过时间
	CheckingPassTime string `json:"checking_pass_time,omitempty" xml:"checking_pass_time,omitempty"`
	// 关单原因
	CloseMemo string `json:"close_memo,omitempty" xml:"close_memo,omitempty"`
	// 改签关单时间
	CloseTime string `json:"close_time,omitempty" xml:"close_time,omitempty"`
	// 申请单改签时间
	GmtCreateTime string `json:"gmt_create_time,omitempty" xml:"gmt_create_time,omitempty"`
	// 改签修改时间
	GmtModifiedTime string `json:"gmt_modified_time,omitempty" xml:"gmt_modified_time,omitempty"`
	// 代理商最晚出票时间
	LatestAgentIssueTime string `json:"latest_agent_issue_time,omitempty" xml:"latest_agent_issue_time,omitempty"`
	// 代理商最迟确认时间
	LatestAgentProcessingTime string `json:"latest_agent_processing_time,omitempty" xml:"latest_agent_processing_time,omitempty"`
	// 买家最迟支付时间
	LatestBuyerPayTime string `json:"latest_buyer_pay_time,omitempty" xml:"latest_buyer_pay_time,omitempty"`
	// 支付成功时间
	PaySuccessTime string `json:"pay_success_time,omitempty" xml:"pay_success_time,omitempty"`
	// 改签意向
	ChangeIntensions string `json:"change_intensions,omitempty" xml:"change_intensions,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 买家ID
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 改签产品信息
	ChangeItem *IeChangeItemVo `json:"change_item,omitempty" xml:"change_item,omitempty"`
	// 改签申请类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿)
	ChangeReasonType int64 `json:"change_reason_type,omitempty" xml:"change_reason_type,omitempty"`
	// 关单类型 关单类型 -1-未关闭,0-买家主动关闭,1-买家支付超时,10-卖家确认处理超时,11-卖家拒绝,99-其他原因
	CloseType int64 `json:"close_type,omitempty" xml:"close_type,omitempty"`
	// 改签单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 改签单业务状态 10－卖家待确认,20-卖家已确认,30-卖家待出票,40-出票超时冻结,50-出票成功,60-验真失败,61-验真成功,70-卖家已拒绝
	OrderBizStatus int64 `json:"order_biz_status,omitempty" xml:"order_biz_status,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 改签订单状态 0－初始状态,1－待支付,2－付款成功,3-改签成功 ,4－改签关闭
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 支付状态 0-初始状态,1-待支付,2-支付成功,3-转交易成功(打款给卖家),4-关闭已支付订单成功,5-关闭未支付订单成功
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 订单总金额
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 改签联系人
	ChangeContactVO *IeChangeContactVo `json:"change_contact_v_o,omitempty" xml:"change_contact_v_o,omitempty"`
}

var poolIeChangeOrderVo = sync.Pool{
	New: func() any {
		return new(IeChangeOrderVo)
	},
}

// GetIeChangeOrderVo() 从对象池中获取IeChangeOrderVo
func GetIeChangeOrderVo() *IeChangeOrderVo {
	return poolIeChangeOrderVo.Get().(*IeChangeOrderVo)
}

// ReleaseIeChangeOrderVo 释放IeChangeOrderVo
func ReleaseIeChangeOrderVo(v *IeChangeOrderVo) {
	v.ChangePassengerVOS = v.ChangePassengerVOS[:0]
	v.ChangeTicketVOS = v.ChangeTicketVOS[:0]
	v.AlipayTradeNO = ""
	v.ApplyTime = ""
	v.BuyerIntensionMemo = ""
	v.CheckingPassTime = ""
	v.CloseMemo = ""
	v.CloseTime = ""
	v.GmtCreateTime = ""
	v.GmtModifiedTime = ""
	v.LatestAgentIssueTime = ""
	v.LatestAgentProcessingTime = ""
	v.LatestBuyerPayTime = ""
	v.PaySuccessTime = ""
	v.ChangeIntensions = ""
	v.AgentId = 0
	v.BuyerId = 0
	v.ChangeItem = nil
	v.ChangeReasonType = 0
	v.CloseType = 0
	v.Id = 0
	v.OrderBizStatus = 0
	v.OrderId = 0
	v.OrderStatus = 0
	v.PayStatus = 0
	v.TotalPrice = 0
	v.ChangeContactVO = nil
	poolIeChangeOrderVo.Put(v)
}
