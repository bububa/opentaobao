package alscmerchant

import (
	"sync"
)

// TicketInfo 结构体
type TicketInfo struct {
	// 凭证ID。核销接口入参
	TicketId string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// 下单时的商品名称
	TicketName string `json:"ticket_name,omitempty" xml:"ticket_name,omitempty"`
	// 凭证券码。仅入参券码对应的凭证会返回该值，其余凭证此属性无返回值
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
	// 凭证金额，单位分。如果是团购，为单份凭证的金额（对应商品售价,不是商品原价,不能作为代金券抵扣金额）；如果是N次次卡，则为N次卡的总金额，N 对应quantity值
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 商户实收金额（抽佣前），单位分。团购和次卡场景，解释同total_amount。公式：凭证金额-实收金额=商户优惠金额
	RealAmount int64 `json:"real_amount,omitempty" xml:"real_amount,omitempty"`
	// 团购场景，固定为1，购买多份与此值无关；次卡场景，例如3次卡，该属性返回3
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 购买时商品的原价（次卡是多次的总原价），单位分
	OriginalPrice int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
}

var poolTicketInfo = sync.Pool{
	New: func() any {
		return new(TicketInfo)
	},
}

// GetTicketInfo() 从对象池中获取TicketInfo
func GetTicketInfo() *TicketInfo {
	return poolTicketInfo.Get().(*TicketInfo)
}

// ReleaseTicketInfo 释放TicketInfo
func ReleaseTicketInfo(v *TicketInfo) {
	v.TicketId = ""
	v.TicketName = ""
	v.TicketCode = ""
	v.TotalAmount = 0
	v.RealAmount = 0
	v.Quantity = 0
	v.OriginalPrice = 0
	poolTicketInfo.Put(v)
}
