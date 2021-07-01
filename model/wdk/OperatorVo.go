package wdk

// OperatorVo 结构体
type OperatorVo struct {
	// 操作人id
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人name
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 操作人type
	OperatorType int64 `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
}
