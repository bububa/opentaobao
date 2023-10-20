package alscmerchant

// AlibabaalscdaodianticketconsultResult 结构体
type AlibabaalscdaodianticketconsultResult struct {
	// 错误码，success=false时有效
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案，success=false时有效
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结构体
	Value *TicketConsultResponse `json:"value,omitempty" xml:"value,omitempty"`
	// 处理结果是否成功，true为成功，false为失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
