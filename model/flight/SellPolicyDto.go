package flight

// SellPolicyDto 结构体
type SellPolicyDto struct {
	// 政策ID
	PolicyId string `json:"policy_id,omitempty" xml:"policy_id,omitempty"`
	// 销售政策备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 政策代码-新后台使用字段
	PolicyCode string `json:"policy_code,omitempty" xml:"policy_code,omitempty"`
	// 政策备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 政策代码-老后台单子使用
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 机票政策类型：0，默认；1，自定义
	PolicyType int64 `json:"policy_type,omitempty" xml:"policy_type,omitempty"`
	// 销售方式：1-3 机+X，4-5返现
	SaleModeCode int64 `json:"sale_mode_code,omitempty" xml:"sale_mode_code,omitempty"`
}
