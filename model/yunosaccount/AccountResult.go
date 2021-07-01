package yunosaccount

// AccountResult 结构体
type AccountResult struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
