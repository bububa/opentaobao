package ticket

import (
	"sync"
)

// TicketSimpleSkuParam 结构体
type TicketSimpleSkuParam struct {
	// 该票种下使用的价格规则
	PriceRules []PriceRule `json:"price_rules,omitempty" xml:"price_rules>price_rule,omitempty"`
	// 门票 票种类型
	TicketType string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// 门票场次（场次门票专用）
	TicketSeason string `json:"ticket_season,omitempty" xml:"ticket_season,omitempty"`
	// 门票区域（场次门票专用）
	TicketArea string `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
}

var poolTicketSimpleSkuParam = sync.Pool{
	New: func() any {
		return new(TicketSimpleSkuParam)
	},
}

// GetTicketSimpleSkuParam() 从对象池中获取TicketSimpleSkuParam
func GetTicketSimpleSkuParam() *TicketSimpleSkuParam {
	return poolTicketSimpleSkuParam.Get().(*TicketSimpleSkuParam)
}

// ReleaseTicketSimpleSkuParam 释放TicketSimpleSkuParam
func ReleaseTicketSimpleSkuParam(v *TicketSimpleSkuParam) {
	v.PriceRules = v.PriceRules[:0]
	v.TicketType = ""
	v.TicketSeason = ""
	v.TicketArea = ""
	poolTicketSimpleSkuParam.Put(v)
}
