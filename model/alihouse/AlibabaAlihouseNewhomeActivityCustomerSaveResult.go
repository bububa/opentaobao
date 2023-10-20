package alihouse

// AlibabaalihousenewhomeactivitycustomersaveResult 结构体
type AlibabaalihousenewhomeactivitycustomersaveResult struct {
	// 处理失败客户集合
	Data []ActivityCustomerErrHandlerResultDto `json:"data,omitempty" xml:"data>activity_customer_err_handler_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
