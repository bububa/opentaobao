package security

// RpAuditComparisonDetailBo 结构体
type RpAuditComparisonDetailBo struct {
	// 分数
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 对比结果
	Value *RpAuditValueBo `json:"value,omitempty" xml:"value,omitempty"`
	// 类型
	ResultType *RpAuditTypeBo `json:"result_type,omitempty" xml:"result_type,omitempty"`
}
