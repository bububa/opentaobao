package product

// TaobaoBanamadpcItemSelectPropApiResult 结构体
type TaobaoBanamadpcItemSelectPropApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 入参类目下入参属性的子属性schema xml
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}
