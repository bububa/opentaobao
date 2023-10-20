package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintasmsnumsendAPIRequest 短信发送 API请求
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
type AlibabaaliqintasmsnumsendAPIRequest struct {
	model.Params
	// 接收号码
	_recNum string
	// 短信模板CODE
	_smsTemplateCode string
	// 公共回传参数
	_extend string
	// 短信签名
	_smsFreeSignName string
	// 短信模板变量，AckNum是变量参数
	_smsParam string
	// 类型，normal：短信
	_smsType string
	// 商家自定义扩展码
	_extendCode string
	// 商家自定义扩展名,例如店铺nick
	_extendName string
}

// NewAlibabaaliqintasmsnumsendRequest 初始化AlibabaaliqintasmsnumsendAPIRequest对象
func NewAlibabaaliqintasmsnumsendRequest() *AlibabaaliqintasmsnumsendAPIRequest {
	return &AlibabaaliqintasmsnumsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqintasmsnumsendAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.ta.sms.num.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqintasmsnumsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqintasmsnumsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecNum is RecNum Setter
// 接收号码
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetRecNum(_recNum string) error {
	r._recNum = _recNum
	r.Set("rec_num", _recNum)
	return nil
}

// GetRecNum RecNum Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetRecNum() string {
	return r._recNum
}

// SetSmsTemplateCode is SmsTemplateCode Setter
// 短信模板CODE
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetSmsTemplateCode(_smsTemplateCode string) error {
	r._smsTemplateCode = _smsTemplateCode
	r.Set("sms_template_code", _smsTemplateCode)
	return nil
}

// GetSmsTemplateCode SmsTemplateCode Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetSmsTemplateCode() string {
	return r._smsTemplateCode
}

// SetExtend is Extend Setter
// 公共回传参数
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetExtend() string {
	return r._extend
}

// SetSmsFreeSignName is SmsFreeSignName Setter
// 短信签名
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetSmsFreeSignName(_smsFreeSignName string) error {
	r._smsFreeSignName = _smsFreeSignName
	r.Set("sms_free_sign_name", _smsFreeSignName)
	return nil
}

// GetSmsFreeSignName SmsFreeSignName Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetSmsFreeSignName() string {
	return r._smsFreeSignName
}

// SetSmsParam is SmsParam Setter
// 短信模板变量，AckNum是变量参数
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// GetSmsParam SmsParam Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetSmsParam() string {
	return r._smsParam
}

// SetSmsType is SmsType Setter
// 类型，normal：短信
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetSmsType(_smsType string) error {
	r._smsType = _smsType
	r.Set("sms_type", _smsType)
	return nil
}

// GetSmsType SmsType Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetSmsType() string {
	return r._smsType
}

// SetExtendCode is ExtendCode Setter
// 商家自定义扩展码
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetExtendCode(_extendCode string) error {
	r._extendCode = _extendCode
	r.Set("extend_code", _extendCode)
	return nil
}

// GetExtendCode ExtendCode Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetExtendCode() string {
	return r._extendCode
}

// SetExtendName is ExtendName Setter
// 商家自定义扩展名,例如店铺nick
func (r *AlibabaaliqintasmsnumsendAPIRequest) SetExtendName(_extendName string) error {
	r._extendName = _extendName
	r.Set("extend_name", _extendName)
	return nil
}

// GetExtendName ExtendName Getter
func (r AlibabaaliqintasmsnumsendAPIRequest) GetExtendName() string {
	return r._extendName
}
