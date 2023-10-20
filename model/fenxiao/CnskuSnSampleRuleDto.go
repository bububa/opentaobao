package fenxiao

import (
	"sync"
)

// CnskuSnSampleRuleDto 结构体
type CnskuSnSampleRuleDto struct {
	// 规则正则表达式
	RuleRegularExpression string `json:"rule_regular_expression,omitempty" xml:"rule_regular_expression,omitempty"`
	// 规则描述
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// 规则对应图面urlurl
	RuleImgUrl string `json:"rule_img_url,omitempty" xml:"rule_img_url,omitempty"`
	// SN截取正则表达式
	SnMgtSubExpression string `json:"sn_mgt_sub_expression,omitempty" xml:"sn_mgt_sub_expression,omitempty"`
	// 规则示例
	RuleSample string `json:"rule_sample,omitempty" xml:"rule_sample,omitempty"`
	// SN截取结束位置
	SnMgtSubEnd int64 `json:"sn_mgt_sub_end,omitempty" xml:"sn_mgt_sub_end,omitempty"`
	// SN截取开始位置
	SnMgtSubStart int64 `json:"sn_mgt_sub_start,omitempty" xml:"sn_mgt_sub_start,omitempty"`
	// SN是否需要截取
	IsSnMgtSub bool `json:"is_sn_mgt_sub,omitempty" xml:"is_sn_mgt_sub,omitempty"`
}

var poolCnskuSnSampleRuleDto = sync.Pool{
	New: func() any {
		return new(CnskuSnSampleRuleDto)
	},
}

// GetCnskuSnSampleRuleDto() 从对象池中获取CnskuSnSampleRuleDto
func GetCnskuSnSampleRuleDto() *CnskuSnSampleRuleDto {
	return poolCnskuSnSampleRuleDto.Get().(*CnskuSnSampleRuleDto)
}

// ReleaseCnskuSnSampleRuleDto 释放CnskuSnSampleRuleDto
func ReleaseCnskuSnSampleRuleDto(v *CnskuSnSampleRuleDto) {
	v.RuleRegularExpression = ""
	v.RuleDesc = ""
	v.RuleImgUrl = ""
	v.SnMgtSubExpression = ""
	v.RuleSample = ""
	v.SnMgtSubEnd = 0
	v.SnMgtSubStart = 0
	v.IsSnMgtSub = false
	poolCnskuSnSampleRuleDto.Put(v)
}
