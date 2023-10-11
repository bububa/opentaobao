package ascp

// CapacityRuleUpdateResponse 结构体
type CapacityRuleUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
