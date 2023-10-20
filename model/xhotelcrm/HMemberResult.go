package xhotelcrm

// HmemberResult 结构体
type HmemberResult struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 信息
	ResultInfo string `json:"result_info,omitempty" xml:"result_info,omitempty"`
	// 数据
	Module *MerchantMemberBindInfoVo `json:"module,omitempty" xml:"module,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
