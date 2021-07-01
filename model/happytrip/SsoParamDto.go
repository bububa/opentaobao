package happytrip

// SsoParamDto 结构体
type SsoParamDto struct {
	// 免登验证ticket
	Ticket string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}
