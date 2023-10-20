package axindata

import (
	"sync"
)

// FscProductRuleApiDto 结构体
type FscProductRuleApiDto struct {
	// FEE_INCLUDE费用包含;FEE_EXCLUDE费用不包含;EXTRA_COST另付费项目;SHOPPING购物说明;BOOKING_RULE预订须知;BREACH_CLAUSE违约条款;REFUND_RULE退改规则;SAFETY_GUIDE安全须知;KIND_TIPS温馨提示;VISA_INFO签证信息;VISA_ATTACHMENT签证附件;OTHER_DESC其他说明;TRAFFIC_DESC交通说明
	RuleType string `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 补充说明内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolFscProductRuleApiDto = sync.Pool{
	New: func() any {
		return new(FscProductRuleApiDto)
	},
}

// GetFscProductRuleApiDto() 从对象池中获取FscProductRuleApiDto
func GetFscProductRuleApiDto() *FscProductRuleApiDto {
	return poolFscProductRuleApiDto.Get().(*FscProductRuleApiDto)
}

// ReleaseFscProductRuleApiDto 释放FscProductRuleApiDto
func ReleaseFscProductRuleApiDto(v *FscProductRuleApiDto) {
	v.RuleType = ""
	v.Content = ""
	poolFscProductRuleApiDto.Put(v)
}
