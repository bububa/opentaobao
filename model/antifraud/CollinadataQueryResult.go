package antifraud

// CollinadataQueryResult 结构体
type CollinadataQueryResult struct {
	// 积分信息.千分制
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// pass,review,reject
	RiskLevel string `json:"risk_level,omitempty" xml:"risk_level,omitempty"`
	// 字符串格式, 关于score生成的描述信息. 本字段可能为空.
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
}
