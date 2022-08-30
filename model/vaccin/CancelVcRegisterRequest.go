package vaccin

// CancelVcRegisterRequest 结构体
type CancelVcRegisterRequest struct {
	// cdc侧的登记单id
	RegisterId string `json:"register_id,omitempty" xml:"register_id,omitempty"`
	// 健康侧的用户支付宝id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 取消登记时间，会校验格式 yyyy-MM-dd HH:mm:ss
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}
