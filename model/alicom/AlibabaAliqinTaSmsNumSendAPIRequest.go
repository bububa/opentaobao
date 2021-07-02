package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaSmsNumSendAPIRequest 短信发送 API请求
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
type AlibabaAliqinTaSmsNumSendAPIRequest struct {
	model.Params
	// 公共回传参数
	_extend string
	// 类型，normal：短信
	_smsType string
	// 短信签名
	_smsFreeSignName string
	// 短信模板变量，AckNum是变量参数
	_smsParam string
	// 接收号码
	_recNum string
	// 短信模板CODE
	_smsTemplateCode string
	// 商家自定义扩展码
	_extendCode string
	// 商家自定义扩展名,例如店铺nick
	_extendName string
}

// NewAlibabaAliqinTaSmsNumSendRequest 初始化AlibabaAliqinTaSmsNumSendAPIRequest对象
func NewAlibabaAliqinTaSmsNumSendRequest() *AlibabaAliqinTaSmsNumSendAPIRequest {
	return &AlibabaAliqinTaSmsNumSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.sms.num.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Extend Setter
// 公共回传参数
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// Get Extend Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtend() string {
	return r._extend
}

// Set is SmsType Setter
// 类型，normal：短信
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsType(_smsType string) error {
	r._smsType = _smsType
	r.Set("sms_type", _smsType)
	return nil
}

// Get SmsType Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsType() string {
	return r._smsType
}

// Set is SmsFreeSignName Setter
// 短信签名
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsFreeSignName(_smsFreeSignName string) error {
	r._smsFreeSignName = _smsFreeSignName
	r.Set("sms_free_sign_name", _smsFreeSignName)
	return nil
}

// Get SmsFreeSignName Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsFreeSignName() string {
	return r._smsFreeSignName
}

// Set is SmsParam Setter
// 短信模板变量，AckNum是变量参数
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// Get SmsParam Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsParam() string {
	return r._smsParam
}

// Set is RecNum Setter
// 接收号码
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// Get RecNum Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetRecNum() string {
	return r._recNum
}

// Set is SmsTemplateCode Setter
// 短信模板CODE
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// Get SmsTemplateCode Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}

// Set is ExtendCode Setter
// 商家自定义扩展码
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtendCode(_extendCode string) error {
	r._extendCode = _extendCode
	r.Set("extend_code", _extendCode)
	return nil
}

// Get ExtendCode Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtendCode() string {
	return r._extendCode
}

// Set is ExtendName Setter
// 商家自定义扩展名,例如店铺nick
func (r *AlibabaAliqinTaSmsNumSendAPIRequest) SetExtendName(_extendName string) error {
	r._extendName = _extendName
	r.Set("extend_name", _extendName)
	return nil
}

// Get ExtendName Getter
func (r AlibabaAliqinTaSmsNumSendAPIRequest) GetExtendName() string {
	return r._extendName
}
