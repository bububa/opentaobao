package security

// RealNameResult 结构体
type RealNameResult struct {
	// checkCode
	CheckCode string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// checkMessage
	CheckMessage string `json:"check_message,omitempty" xml:"check_message,omitempty"`
	// match
	Match bool `json:"match,omitempty" xml:"match,omitempty"`
}
