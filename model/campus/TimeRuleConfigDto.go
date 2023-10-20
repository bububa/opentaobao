package campus

import (
	"sync"
)

// TimeRuleConfigDto 结构体
type TimeRuleConfigDto struct {
	// 时间规则集合
	TimeRuleList []TimeRuleConfigDto `json:"time_rule_list,omitempty" xml:"time_rule_list>time_rule_config_dto,omitempty"`
	// 失效条件
	ForbidTemplateList []CalendarTemplateDto `json:"forbid_template_list,omitempty" xml:"forbid_template_list>calendar_template_dto,omitempty"`
	// 通行日期
	AllowTemplateList []CalendarTemplateDto `json:"allow_template_list,omitempty" xml:"allow_template_list>calendar_template_dto,omitempty"`
	// 规则名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 周
	Week string `json:"week,omitempty" xml:"week,omitempty"`
	// 时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 时间规则ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTimeRuleConfigDto = sync.Pool{
	New: func() any {
		return new(TimeRuleConfigDto)
	},
}

// GetTimeRuleConfigDto() 从对象池中获取TimeRuleConfigDto
func GetTimeRuleConfigDto() *TimeRuleConfigDto {
	return poolTimeRuleConfigDto.Get().(*TimeRuleConfigDto)
}

// ReleaseTimeRuleConfigDto 释放TimeRuleConfigDto
func ReleaseTimeRuleConfigDto(v *TimeRuleConfigDto) {
	v.TimeRuleList = v.TimeRuleList[:0]
	v.ForbidTemplateList = v.ForbidTemplateList[:0]
	v.AllowTemplateList = v.AllowTemplateList[:0]
	v.Name = ""
	v.Week = ""
	v.Time = ""
	v.Id = 0
	poolTimeRuleConfigDto.Put(v)
}
