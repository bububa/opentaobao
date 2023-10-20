package promotion

import (
	"sync"
)

// ShowRuleDto 结构体
type ShowRuleDto struct {
	// 规则扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 规则类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 规则是否通过
	Passed bool `json:"passed,omitempty" xml:"passed,omitempty"`
}

var poolShowRuleDto = sync.Pool{
	New: func() any {
		return new(ShowRuleDto)
	},
}

// GetShowRuleDto() 从对象池中获取ShowRuleDto
func GetShowRuleDto() *ShowRuleDto {
	return poolShowRuleDto.Get().(*ShowRuleDto)
}

// ReleaseShowRuleDto 释放ShowRuleDto
func ReleaseShowRuleDto(v *ShowRuleDto) {
	v.Feature = ""
	v.Type = ""
	v.Passed = false
	poolShowRuleDto.Put(v)
}
