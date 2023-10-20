package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtsmscodesendAPIRequest 喵零发送短信 API请求
// tmall.nrt.sms.code.send
//
// 喵零发送短信
type TmallnrtsmscodesendAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 业务身份
	_bizCode string
	// 1:喵零入会验证码
	_type string
	// 短信参数
	_smsParam string
}

// NewTmallnrtsmscodesendRequest 初始化TmallnrtsmscodesendAPIRequest对象
func NewTmallnrtsmscodesendRequest() *TmallnrtsmscodesendAPIRequest {
	return &TmallnrtsmscodesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtsmscodesendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.sms.code.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtsmscodesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtsmscodesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *TmallnrtsmscodesendAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TmallnrtsmscodesendAPIRequest) GetPhone() string {
	return r._phone
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallnrtsmscodesendAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallnrtsmscodesendAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetType is Type Setter
// 1:喵零入会验证码
func (r *TmallnrtsmscodesendAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallnrtsmscodesendAPIRequest) GetType() string {
	return r._type
}

// SetSmsParam is SmsParam Setter
// 短信参数
func (r *TmallnrtsmscodesendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// GetSmsParam SmsParam Getter
func (r TmallnrtsmscodesendAPIRequest) GetSmsParam() string {
	return r._smsParam
}
