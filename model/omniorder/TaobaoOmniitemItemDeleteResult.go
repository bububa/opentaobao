package omniorder

// TaobaoOmniitemItemDeleteResult 结构体
type TaobaoOmniitemItemDeleteResult struct {
	// 错误码，0表示成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *ItemDeleteResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
