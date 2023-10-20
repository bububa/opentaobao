package tmallnr

// TmallNrtCertificateQueryResult 结构体
type TmallNrtCertificateQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// model
	Model *PageData `json:"model,omitempty" xml:"model,omitempty"`
}
