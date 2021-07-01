package alimsg

// SendTemplateMsgResponse 结构体
type SendTemplateMsgResponse struct {
	// 32位的uuid,用来追踪一条消息
	MessageTraceId string `json:"message_trace_id,omitempty" xml:"message_trace_id,omitempty"`
}
