package alihouse

// ActivityCustomerErrHandlerResultDto 结构体
type ActivityCustomerErrHandlerResultDto struct {
	// 外部客户id
	OuterCustomerId string `json:"outer_customer_id,omitempty" xml:"outer_customer_id,omitempty"`
	// 错误原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 是否删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}
