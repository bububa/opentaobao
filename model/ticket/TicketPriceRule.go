package ticket

import (
	"sync"
)

// TicketPriceRule 结构体
type TicketPriceRule struct {
	// 必填，该票种下使用的价格规则。
	PriceRules []PriceRule `json:"price_rules,omitempty" xml:"price_rules>price_rule,omitempty"`
	// 必填，门票 票种类型
	TicketType string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// 可选，门票场次（场次门票专用，对于场次门票必选）
	TicketSeason string `json:"ticket_season,omitempty" xml:"ticket_season,omitempty"`
	// 可选，门票区域（场次门票专用，对于场次门票必选）
	TicketArea string `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
}

var poolTicketPriceRule = sync.Pool{
	New: func() any {
		return new(TicketPriceRule)
	},
}

// GetTicketPriceRule() 从对象池中获取TicketPriceRule
func GetTicketPriceRule() *TicketPriceRule {
	return poolTicketPriceRule.Get().(*TicketPriceRule)
}

// ReleaseTicketPriceRule 释放TicketPriceRule
func ReleaseTicketPriceRule(v *TicketPriceRule) {
	v.PriceRules = v.PriceRules[:0]
	v.TicketType = ""
	v.TicketSeason = ""
	v.TicketArea = ""
	poolTicketPriceRule.Put(v)
}
