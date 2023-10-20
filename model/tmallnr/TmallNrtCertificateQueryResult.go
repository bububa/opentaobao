package tmallnr

// TmallnrtcertificatequeryResult 结构体
type TmallnrtcertificatequeryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// model
	Model *PageData `json:"model,omitempty" xml:"model,omitempty"`
}
