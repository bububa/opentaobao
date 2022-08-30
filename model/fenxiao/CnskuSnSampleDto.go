package fenxiao

// CnskuSnSampleDto 结构体
type CnskuSnSampleDto struct {
	// SN示例
	SampleRuleList []CnskuSnSampleRuleDto `json:"sample_rule_list,omitempty" xml:"sample_rule_list>cnsku_sn_sample_rule_dto,omitempty"`
	// sn示例顺序
	SnSeq string `json:"sn_seq,omitempty" xml:"sn_seq,omitempty"`
	// sn示例顺序
	SampleDesc string `json:"sample_desc,omitempty" xml:"sample_desc,omitempty"`
}
