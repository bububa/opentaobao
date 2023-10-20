package ticket

import (
	"sync"
)

// TopTicketRuleResult 结构体
type TopTicketRuleResult struct {
	// 规则名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商家景点编码
	OutScenicId string `json:"out_scenic_id,omitempty" xml:"out_scenic_id,omitempty"`
	// 商家规则编码
	OutRuleId string `json:"out_rule_id,omitempty" xml:"out_rule_id,omitempty"`
	// 规则主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 阿里景点id
	AliScenicId int64 `json:"ali_scenic_id,omitempty" xml:"ali_scenic_id,omitempty"`
}

var poolTopTicketRuleResult = sync.Pool{
	New: func() any {
		return new(TopTicketRuleResult)
	},
}

// GetTopTicketRuleResult() 从对象池中获取TopTicketRuleResult
func GetTopTicketRuleResult() *TopTicketRuleResult {
	return poolTopTicketRuleResult.Get().(*TopTicketRuleResult)
}

// ReleaseTopTicketRuleResult 释放TopTicketRuleResult
func ReleaseTopTicketRuleResult(v *TopTicketRuleResult) {
	v.Name = ""
	v.OutScenicId = ""
	v.OutRuleId = ""
	v.Id = 0
	v.AliScenicId = 0
	poolTopTicketRuleResult.Put(v)
}
