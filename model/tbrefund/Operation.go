package tbrefund

// Operation 结构体
type Operation struct {
	// 操作编码
	OperationCode string `json:"operation_code,omitempty" xml:"operation_code,omitempty"`
	// 操作提示文案
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
}
