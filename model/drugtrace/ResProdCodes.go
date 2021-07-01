package drugtrace

// ResProdCodes 结构体
type ResProdCodes struct {
	// 资源码
	ResCodeList []ResCode `json:"res_code_list,omitempty" xml:"res_code_list>res_code,omitempty"`
}
