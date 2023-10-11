package logistic

// TmsPhoneCallInfoRequest 结构体
type TmsPhoneCallInfoRequest struct {
	// 电联属性，phoneCallInfo有值时，必需。SMART 智能云呼; NORMAL, 普通
	PhoneCallType string `json:"phone_call_type,omitempty" xml:"phone_call_type,omitempty"`
	// 挂断状态，phoneCallInfo有值时，必需。CUSTOMER_HAND_UP,消费者挂断；SENDER_HANG_UP,小件员挂断
	HangUpType string `json:"hang_up_type,omitempty" xml:"hang_up_type,omitempty"`
	// 电联语音文件，用于客诉场景定责使用
	PhoneChatUrl string `json:"phone_chat_url,omitempty" xml:"phone_chat_url,omitempty"`
	// 电联用户后，用户确认需要进行改派时，操作改派的时间
	PhoneCallDemandChangeTime string `json:"phone_call_demand_change_time,omitempty" xml:"phone_call_demand_change_time,omitempty"`
	// 接通时间
	ConnectTime string `json:"connect_time,omitempty" xml:"connect_time,omitempty"`
	// 消费者电联需求描述
	PhoneCallResultRemark string `json:"phone_call_result_remark,omitempty" xml:"phone_call_result_remark,omitempty"`
	// 通话状态，phoneCallInfo有值时，必需。NOT _CONNECT, 未拨通；CONNECTED，已接通； NOBODY_ANSWER,无人接听
	ConnectionStatus string `json:"connection_status,omitempty" xml:"connection_status,omitempty"`
	// 拨打时间。phoneCallInfo有值时，必需
	CallTime string `json:"call_time,omitempty" xml:"call_time,omitempty"`
	// 挂断时间
	HangUpTime string `json:"hang_up_time,omitempty" xml:"hang_up_time,omitempty"`
}
