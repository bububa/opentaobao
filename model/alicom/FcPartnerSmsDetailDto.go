package alicom

// FcPartnerSmsDetailDTO 
type FcPartnerSmsDetailDTO struct {
    // 公共回传参数
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 短信接收号码
    RecNum   string `json:"rec_num,omitempty" xml:"rec_num,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 模板编码
    SmsCode   string `json:"sms_code,omitempty" xml:"sms_code,omitempty"`
    // 短信内容
    SmsContent   string `json:"sms_content,omitempty" xml:"sms_content,omitempty"`
    // 短信接收时间
    SmsReceiverTime   string `json:"sms_receiver_time,omitempty" xml:"sms_receiver_time,omitempty"`
    // 短信发送时间
    SmsSendTime   string `json:"sms_send_time,omitempty" xml:"sms_send_time,omitempty"`
    // 发送状态 1：等待回执，2：发送失败，3：发送成功
    SmsStatus   int64 `json:"sms_status,omitempty" xml:"sms_status,omitempty"`
}
