package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtSmsCodeSendAPIRequest 喵零发送短信 API请求
// tmall.nrt.sms.code.send
//
// 喵零发送短信
type TmallNrtSmsCodeSendAPIRequest struct {
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

// NewTmallNrtSmsCodeSendRequest 初始化TmallNrtSmsCodeSendAPIRequest对象
func NewTmallNrtSmsCodeSendRequest() *TmallNrtSmsCodeSendAPIRequest {
	return &TmallNrtSmsCodeSendAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtSmsCodeSendAPIRequest) Reset() {
	r._phone = ""
	r._bizCode = ""
	r._type = ""
	r._smsParam = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtSmsCodeSendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.sms.code.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtSmsCodeSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtSmsCodeSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *TmallNrtSmsCodeSendAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TmallNrtSmsCodeSendAPIRequest) GetPhone() string {
	return r._phone
}

// SetBizCode is BizCode Setter
// 业务身份
func (r *TmallNrtSmsCodeSendAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TmallNrtSmsCodeSendAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetType is Type Setter
// 1:喵零入会验证码
func (r *TmallNrtSmsCodeSendAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallNrtSmsCodeSendAPIRequest) GetType() string {
	return r._type
}

// SetSmsParam is SmsParam Setter
// 短信参数
func (r *TmallNrtSmsCodeSendAPIRequest) SetSmsParam(_smsParam string) error {
	r._smsParam = _smsParam
	r.Set("sms_param", _smsParam)
	return nil
}

// GetSmsParam SmsParam Getter
func (r TmallNrtSmsCodeSendAPIRequest) GetSmsParam() string {
	return r._smsParam
}

var poolTmallNrtSmsCodeSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtSmsCodeSendRequest()
	},
}

// GetTmallNrtSmsCodeSendRequest 从 sync.Pool 获取 TmallNrtSmsCodeSendAPIRequest
func GetTmallNrtSmsCodeSendAPIRequest() *TmallNrtSmsCodeSendAPIRequest {
	return poolTmallNrtSmsCodeSendAPIRequest.Get().(*TmallNrtSmsCodeSendAPIRequest)
}

// ReleaseTmallNrtSmsCodeSendAPIRequest 将 TmallNrtSmsCodeSendAPIRequest 放入 sync.Pool
func ReleaseTmallNrtSmsCodeSendAPIRequest(v *TmallNrtSmsCodeSendAPIRequest) {
	v.Reset()
	poolTmallNrtSmsCodeSendAPIRequest.Put(v)
}
