package aliqin

// AlibabaaliqinfciotqrycardResult 结构体
type AlibabaaliqinfciotqrycardResult struct {
	// model
	Models []AlibabaaliqinfciotqrycardModel `json:"models,omitempty" xml:"models>alibabaaliqinfciotqrycard_model,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// true返回成功，false返回失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
