package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamedeleteAPIRequest 淘宝短信签名删除 API请求
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
type TaobaojstsmssignnamedeleteAPIRequest struct {
	model.Params
	// 删除签名入参
	_deleteSmsSignRequest *TopDeleteSmsSignRequest
}

// NewTaobaojstsmssignnamedeleteRequest 初始化TaobaojstsmssignnamedeleteAPIRequest对象
func NewTaobaojstsmssignnamedeleteRequest() *TaobaojstsmssignnamedeleteAPIRequest {
	return &TaobaojstsmssignnamedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmssignnamedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmssignnamedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmssignnamedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteSmsSignRequest is DeleteSmsSignRequest Setter
// 删除签名入参
func (r *TaobaojstsmssignnamedeleteAPIRequest) SetDeleteSmsSignRequest(_deleteSmsSignRequest *TopDeleteSmsSignRequest) error {
	r._deleteSmsSignRequest = _deleteSmsSignRequest
	r.Set("delete_sms_sign_request", _deleteSmsSignRequest)
	return nil
}

// GetDeleteSmsSignRequest DeleteSmsSignRequest Getter
func (r TaobaojstsmssignnamedeleteAPIRequest) GetDeleteSmsSignRequest() *TopDeleteSmsSignRequest {
	return r._deleteSmsSignRequest
}
