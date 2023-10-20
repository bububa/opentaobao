package flight

import (
	"sync"
)

// TicketingDetailDto 结构体
type TicketingDetailDto struct {
	// 出票对象
	IssueList []IssueList `json:"issue_list,omitempty" xml:"issue_list>issue_list,omitempty"`
	// 订单标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 行李
	BaggageList []BaggageDto `json:"baggage_list,omitempty" xml:"baggage_list>baggage_dto,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 出票时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// sla
	Sla string `json:"sla,omitempty" xml:"sla,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 退改签规则
	RefundModifyRule string `json:"refund_modify_rule,omitempty" xml:"refund_modify_rule,omitempty"`
	// 意向单id
	IntentionId string `json:"intention_id,omitempty" xml:"intention_id,omitempty"`
	// 支付成功后30分钟内出票
	SlaDesc string `json:"sla_desc,omitempty" xml:"sla_desc,omitempty"`
	// xxx
	CorrelationOutId string `json:"correlation_out_id,omitempty" xml:"correlation_out_id,omitempty"`
	// 催出后服务时效
	UrgeSla string `json:"urge_sla,omitempty" xml:"urge_sla,omitempty"`
	// 2023-05-26 20:13:08后存在订单被取消的风险
	UrgeSlaDesc string `json:"urge_sla_desc,omitempty" xml:"urge_sla_desc,omitempty"`
	// 店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 佣金
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 订单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	//    1:&#34;单程&#34;,     2:&#34;往返&#34;,     3:&#34;多程&#34;
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

var poolTicketingDetailDto = sync.Pool{
	New: func() any {
		return new(TicketingDetailDto)
	},
}

// GetTicketingDetailDto() 从对象池中获取TicketingDetailDto
func GetTicketingDetailDto() *TicketingDetailDto {
	return poolTicketingDetailDto.Get().(*TicketingDetailDto)
}

// ReleaseTicketingDetailDto 释放TicketingDetailDto
func ReleaseTicketingDetailDto(v *TicketingDetailDto) {
	v.IssueList = v.IssueList[:0]
	v.Tags = v.Tags[:0]
	v.BaggageList = v.BaggageList[:0]
	v.OrderId = ""
	v.PayTime = ""
	v.IssueTime = ""
	v.Sla = ""
	v.Currency = ""
	v.RefundModifyRule = ""
	v.IntentionId = ""
	v.SlaDesc = ""
	v.CorrelationOutId = ""
	v.UrgeSla = ""
	v.UrgeSlaDesc = ""
	v.AgentId = 0
	v.DomesticIntl = 0
	v.Commission = 0
	v.Status = 0
	v.TripType = 0
	poolTicketingDetailDto.Put(v)
}
