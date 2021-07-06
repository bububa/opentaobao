package alicom

// CommonRtnDo 结构体
type CommonRtnDo struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
