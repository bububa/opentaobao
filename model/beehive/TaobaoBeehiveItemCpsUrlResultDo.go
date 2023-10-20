package beehive

// TaobaobeehiveitemcpsurlResultDo 结构体
type TaobaobeehiveitemcpsurlResultDo struct {
	// 商品id和对应的url map
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
