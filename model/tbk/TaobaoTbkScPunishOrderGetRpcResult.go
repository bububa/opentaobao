package tbk

// TaobaotbkscpunishordergetRpcResult 结构体
type TaobaotbkscpunishordergetRpcResult struct {
	// 业务出错的描述
	Bizerrordesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 执行结果
	Resultmsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 业务出错的状态码
	Bizerrorcode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 执行结果状态码
	Resultcode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
