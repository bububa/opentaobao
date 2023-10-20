package fenxiao

import (
	"sync"
)

// CnskuSnSampleDto 结构体
type CnskuSnSampleDto struct {
	// SN示例
	SampleRuleList []CnskuSnSampleRuleDto `json:"sample_rule_list,omitempty" xml:"sample_rule_list>cnsku_sn_sample_rule_dto,omitempty"`
	// sn示例顺序
	SnSeq string `json:"sn_seq,omitempty" xml:"sn_seq,omitempty"`
	// sn示例顺序
	SampleDesc string `json:"sample_desc,omitempty" xml:"sample_desc,omitempty"`
}

var poolCnskuSnSampleDto = sync.Pool{
	New: func() any {
		return new(CnskuSnSampleDto)
	},
}

// GetCnskuSnSampleDto() 从对象池中获取CnskuSnSampleDto
func GetCnskuSnSampleDto() *CnskuSnSampleDto {
	return poolCnskuSnSampleDto.Get().(*CnskuSnSampleDto)
}

// ReleaseCnskuSnSampleDto 释放CnskuSnSampleDto
func ReleaseCnskuSnSampleDto(v *CnskuSnSampleDto) {
	v.SampleRuleList = v.SampleRuleList[:0]
	v.SnSeq = ""
	v.SampleDesc = ""
	poolCnskuSnSampleDto.Put(v)
}
