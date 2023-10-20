package tmallnr

// TmallnrtcertificatesendResult 结构体
type TmallnrtcertificatesendResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// model
	Model *NrtCertificateInstanceResponseDto `json:"model,omitempty" xml:"model,omitempty"`
}
