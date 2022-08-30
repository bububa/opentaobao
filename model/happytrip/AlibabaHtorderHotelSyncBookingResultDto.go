package happytrip

// AlibabaHtorderHotelSyncBookingResultDto 结构体
type AlibabaHtorderHotelSyncBookingResultDto struct {
	// 错误码
	ErrNo string `json:"err_no,omitempty" xml:"err_no,omitempty"`
	// 错误信息
	ErrInfo string `json:"err_info,omitempty" xml:"err_info,omitempty"`
	// 栈信息
	StackTrace string `json:"stack_trace,omitempty" xml:"stack_trace,omitempty"`
	// 每个订单的信息发送结果
	Content *SyncHotelBookingDataResponseDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
