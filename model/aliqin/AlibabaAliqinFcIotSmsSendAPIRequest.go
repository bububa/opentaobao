package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotSmsSendAPIRequest 物联网短信发送 API请求
// alibaba.aliqin.fc.iot.sms.send
//
// 发送物联网短信，只允许使用物联网短信模板
type AlibabaAliqinFcIotSmsSendAPIRequest struct {
	model.Params
	// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	_extend string
	// 短信类型，传入值请填写normal
	_smsType string
	// 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
	_smsParam string
	// 短信接收号码。
	_recNum string
	// 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
	_smsTemplateCode string
}

// NewAlibabaAliqinFcIotSmsSendRequest 初始化AlibabaAliqinFcIotSmsSendAPIRequest对象
func NewAlibabaAliqinFcIotSmsSendRequest() *AlibabaAliqinFcIotSmsSendAPIRequest {
	return &AlibabaAliqinFcIotSmsSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.sms.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtend is Extend Setter
// 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
func (r *AlibabaAliqinFcIotSmsSendAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetExtend() string {
	return r._extend
}

// SetSmsType is SmsType Setter
// 短信类型，传入值请填写normal
func (r *AlibabaAliqinFcIotSmsSendAPIRequest) SetSmsType(_smsType string) error {
	r._smsType = _smsType
	r.Set("sms_type", _smsType)
	return nil
}

// GetSmsType SmsType Getter
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetSmsType() string {
	return r._smsType
}

// SetSmsParam is SmsParam Setter
// 短信模板变量，传参规则{&#34;key&#34;:&#34;value&#34;}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{&#34;code&#34;:&#34;1234&#34;,&#34;product&#34;:&#34;alidayu&#34;}
func (r *AlibabaAliqinFcIotSmsSendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// GetSmsParam SmsParam Getter
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetSmsParam() string {
	return r._smsParam
}

// SetRecNum is RecNum Setter
// 短信接收号码。
func (r *AlibabaAliqinFcIotSmsSendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetSmsTemplateCode is SmsTemplateCode Setter
// 短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
func (r *AlibabaAliqinFcIotSmsSendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// GetSmsTemplateCode SmsTemplateCode Getter
func (r AlibabaAliqinFcIotSmsSendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}
