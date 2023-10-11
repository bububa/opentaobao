package train

// FreeChildrenTicketDealRq 结构体
type FreeChildrenTicketDealRq struct {
	// 失败文案
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 唯一id
	ApplyNo int64 `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
	// 失败code
	FailCode int64 `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
