package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamequeryAPIRequest 淘宝短信签名查询 API请求
// taobao.jst.sms.signname.query
//
// 淘宝短信签名查询
type TaobaojstsmssignnamequeryAPIRequest struct {
	model.Params
	// 签名查询的入参
	_querySmsSignRequest *TopQuerySmsSignRequest
}

// NewTaobaojstsmssignnamequeryRequest 初始化TaobaojstsmssignnamequeryAPIRequest对象
func NewTaobaojstsmssignnamequeryRequest() *TaobaojstsmssignnamequeryAPIRequest {
	return &TaobaojstsmssignnamequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmssignnamequeryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmssignnamequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmssignnamequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuerySmsSignRequest is QuerySmsSignRequest Setter
// 签名查询的入参
func (r *TaobaojstsmssignnamequeryAPIRequest) SetQuerySmsSignRequest(_querySmsSignRequest *TopQuerySmsSignRequest) error {
	r._querySmsSignRequest = _querySmsSignRequest
	r.Set("query_sms_sign_request", _querySmsSignRequest)
	return nil
}

// GetQuerySmsSignRequest QuerySmsSignRequest Getter
func (r TaobaojstsmssignnamequeryAPIRequest) GetQuerySmsSignRequest() *TopQuerySmsSignRequest {
	return r._querySmsSignRequest
}
