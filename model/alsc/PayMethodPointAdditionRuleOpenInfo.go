package alsc

// PayMethodPointAdditionRuleOpenInfo 结构体
type PayMethodPointAdditionRuleOpenInfo struct {
	// 是否开启
	Enable bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 方法ID
	MethodId string `json:"method_id,omitempty" xml:"method_id,omitempty"`
	// 方法名称
	MethodName string `json:"method_name,omitempty" xml:"method_name,omitempty"`
}
