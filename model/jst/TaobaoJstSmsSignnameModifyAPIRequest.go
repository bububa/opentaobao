package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamemodifyAPIRequest 淘宝短信签名修改 API请求
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
type TaobaojstsmssignnamemodifyAPIRequest struct {
	model.Params
	// 修改签名入参
	_modifySmsSignRequest *TopModifySmsSignRequest
}

// NewTaobaojstsmssignnamemodifyRequest 初始化TaobaojstsmssignnamemodifyAPIRequest对象
func NewTaobaojstsmssignnamemodifyRequest() *TaobaojstsmssignnamemodifyAPIRequest {
	return &TaobaojstsmssignnamemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmssignnamemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmssignnamemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmssignnamemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifySmsSignRequest is ModifySmsSignRequest Setter
// 修改签名入参
func (r *TaobaojstsmssignnamemodifyAPIRequest) SetModifySmsSignRequest(_modifySmsSignRequest *TopModifySmsSignRequest) error {
	r._modifySmsSignRequest = _modifySmsSignRequest
	r.Set("modify_sms_sign_request", _modifySmsSignRequest)
	return nil
}

// GetModifySmsSignRequest ModifySmsSignRequest Getter
func (r TaobaojstsmssignnamemodifyAPIRequest) GetModifySmsSignRequest() *TopModifySmsSignRequest {
	return r._modifySmsSignRequest
}
