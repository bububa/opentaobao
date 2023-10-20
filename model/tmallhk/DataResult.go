package tmallhk

// DataResult 结构体
type DataResult struct {
	// 参数code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 参数msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// obj
	Obj *CciccheckCodeDo `json:"obj,omitempty" xml:"obj,omitempty"`
	// 是否正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
