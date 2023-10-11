package alicom

// Subs 结构体
type Subs struct {
	// 顺振参数
	SequenceCalls []SequenceCalls `json:"sequence_calls,omitempty" xml:"sequence_calls>sequence_calls,omitempty"`
	// 被叫号码
	CalledNo string `json:"called_no,omitempty" xml:"called_no,omitempty"`
	// 短信通道方式SMS_INTERCEPT(拦截推送阿里)，SMS_NORMAL_SEND(正常现网下发)，SMS_INTERCEPT_SEND(截取并转发)，SMS_DROP(拦截丢弃)，
	SmsChannel string `json:"sms_channel,omitempty" xml:"sms_channel,omitempty"`
	// 呼叫类型MASTER(A-&gt;X-&gt;B), CALLED(B-&gt;X-&gt;A), SMS_SENDER, SMS_RECEIVER
	CallType string `json:"call_type,omitempty" xml:"call_type,omitempty"`
	// 录音类型，mp3/wav
	RecType string `json:"rec_type,omitempty" xml:"rec_type,omitempty"`
	// 订购关系ID；目前字符串长度为16位，建议预留32位
	SubsId string `json:"subs_id,omitempty" xml:"subs_id,omitempty"`
	// 被叫显号
	CalledDisplayNo string `json:"called_display_no,omitempty" xml:"called_display_no,omitempty"`
	// 录音模式，1：仅录制通话录音、2：放音录音+通话录音
	RecordMode string `json:"record_mode,omitempty" xml:"record_mode,omitempty"`
	// 是否需要优先下载录音，0：否、1：是
	FastRecord string `json:"fast_record,omitempty" xml:"fast_record,omitempty"`
	// 主叫媒体流推送地址
	WsAddr string `json:"ws_addr,omitempty" xml:"ws_addr,omitempty"`
	// 被叫媒体流推送地址
	WsAddrCalled string `json:"ws_addr_called,omitempty" xml:"ws_addr_called,omitempty"`
	// 外部回传字段
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 实时媒体类型1 彩铃  2 通话  3 彩铃和通话
	RtpType string `json:"rtp_type,omitempty" xml:"rtp_type,omitempty"`
	// 格式定义
	ContentFormat string `json:"content_format,omitempty" xml:"content_format,omitempty"`
	// 顺振超时时间
	SequenceTimeout int64 `json:"sequence_timeout,omitempty" xml:"sequence_timeout,omitempty"`
	// 是否开启铃音检测 0：不开启 1：开启
	RrdsControl int64 `json:"rrds_control,omitempty" xml:"rrds_control,omitempty"`
	// 挂机IVR参数
	EndCallIvr *EndCallIvr `json:"end_call_ivr,omitempty" xml:"end_call_ivr,omitempty"`
	// 顺振规则
	SequenceCallRule *SequenceCallRule `json:"sequence_call_rule,omitempty" xml:"sequence_call_rule,omitempty"`
	// 是否需要状态报告，0：不需要，1：需要
	StatusReport int64 `json:"status_report,omitempty" xml:"status_report,omitempty"`
	// 是否需要录音
	NeedRecord bool `json:"need_record,omitempty" xml:"need_record,omitempty"`
	// 是否实时媒体
	NeedRealtimeMedia bool `json:"need_realtime_media,omitempty" xml:"need_realtime_media,omitempty"`
}
