package wdk

// UtmsResult 结构体
type UtmsResult struct {
	// list
	List []BomProcessDto `json:"list,omitempty" xml:"list>bom_process_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
