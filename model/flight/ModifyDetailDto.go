package flight

import (
	"sync"
)

// ModifyDetailDto 结构体
type ModifyDetailDto struct {
	// 改签数据
	ChangeList []ChangeList `json:"change_list,omitempty" xml:"change_list>change_list,omitempty"`
	// 订单标签:REASSURING_TICKET:安心票
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 申请原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 关联飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// sla
	Sla string `json:"sla,omitempty" xml:"sla,omitempty"`
	// 币种:CNY:人民币, USD:美元, HKD:港元
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 拒绝原因
	RefuseReason string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// xxxx
	CorrelationOrderId string `json:"correlation_order_id,omitempty" xml:"correlation_order_id,omitempty"`
	// 是否自愿:0:非自愿，1:自愿
	Voluntary int64 `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
	// 国际国内标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 改签单状态: 1:待回填, 2:待用户支付, 3:待出票, 4:已完成, 5:已拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 佣金;单位:分
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 0:&#34;原路退回&#34;,1:&#34;退银行卡&#34;,2:&#34;原路退回+退银行卡&#34;
	RefundWayType int64 `json:"refund_way_type,omitempty" xml:"refund_way_type,omitempty"`
	//    1:&#34;单程&#34;,     2:&#34;往返&#34;,     3:&#34;多程&#34;
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

var poolModifyDetailDto = sync.Pool{
	New: func() any {
		return new(ModifyDetailDto)
	},
}

// GetModifyDetailDto() 从对象池中获取ModifyDetailDto
func GetModifyDetailDto() *ModifyDetailDto {
	return poolModifyDetailDto.Get().(*ModifyDetailDto)
}

// ReleaseModifyDetailDto 释放ModifyDetailDto
func ReleaseModifyDetailDto(v *ModifyDetailDto) {
	v.ChangeList = v.ChangeList[:0]
	v.Tags = v.Tags[:0]
	v.ApplyId = ""
	v.Reason = ""
	v.OrderId = ""
	v.PayTime = ""
	v.Sla = ""
	v.Currency = ""
	v.ApplyTime = ""
	v.RefuseReason = ""
	v.CorrelationOrderId = ""
	v.Voluntary = 0
	v.DomesticIntl = 0
	v.Status = 0
	v.Commission = 0
	v.RefundWayType = 0
	v.TripType = 0
	poolModifyDetailDto.Put(v)
}
