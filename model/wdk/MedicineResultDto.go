package wdk

// MedicineResultDto 结构体
type MedicineResultDto struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
