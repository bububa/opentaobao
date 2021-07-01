package jst

// SendMessageRequest 结构体
type SendMessageRequest struct {
	// 拓展Name
	ExtendName string `json:"extend_name,omitempty" xml:"extend_name,omitempty"`
	// 拓展Code
	ExtendCode string `json:"extend_code,omitempty" xml:"extend_code,omitempty"`
	// 短信签名
	SmsFreeSignName string `json:"sms_free_sign_name,omitempty" xml:"sms_free_sign_name,omitempty"`
	// 渠道类型
	ChannelType string `json:"channel_type,omitempty" xml:"channel_type,omitempty"`
	// 短信模版Code
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 短信内容参数, ${url}会被入参url的值替换掉
	Params string `json:"params,omitempty" xml:"params,omitempty"`
	// 商品H5详情页，如果不传则没有短信效果数据
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 拓展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 手机号码
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 短信人群tag
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// 短信批次号
	BatchNumber string `json:"batch_number,omitempty" xml:"batch_number,omitempty"`
	// 短信类型
	SmsType string `json:"sms_type,omitempty" xml:"sms_type,omitempty"`
	// 标记字段
	ToolFlag string `json:"tool_flag,omitempty" xml:"tool_flag,omitempty"`
}
