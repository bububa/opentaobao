package promotion

// AlibababenefitqueryResult 结构体
type AlibababenefitqueryResult struct {
	// datas
	Datas []OrightDto `json:"datas,omitempty" xml:"datas>oright_dto,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
