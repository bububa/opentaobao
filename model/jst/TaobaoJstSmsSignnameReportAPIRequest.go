package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamereportAPIRequest 服务商存量短信签名上传 API请求
// taobao.jst.sms.signname.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的签名数据，确保签名数据正确，否则会导致短信发送失败！！！
type TaobaojstsmssignnamereportAPIRequest struct {
	model.Params
	// 上传签名入参
	_smsSignNameRequest *SmsSignNameRequest
}

// NewTaobaojstsmssignnamereportRequest 初始化TaobaojstsmssignnamereportAPIRequest对象
func NewTaobaojstsmssignnamereportRequest() *TaobaojstsmssignnamereportAPIRequest {
	return &TaobaojstsmssignnamereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmssignnamereportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmssignnamereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmssignnamereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSmsSignNameRequest is SmsSignNameRequest Setter
// 上传签名入参
func (r *TaobaojstsmssignnamereportAPIRequest) SetSmsSignNameRequest(_smsSignNameRequest *SmsSignNameRequest) error {
	r._smsSignNameRequest = _smsSignNameRequest
	r.Set("sms_sign_name_request", _smsSignNameRequest)
	return nil
}

// GetSmsSignNameRequest SmsSignNameRequest Getter
func (r TaobaojstsmssignnamereportAPIRequest) GetSmsSignNameRequest() *SmsSignNameRequest {
	return r._smsSignNameRequest
}
