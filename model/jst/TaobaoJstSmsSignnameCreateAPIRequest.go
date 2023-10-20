package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamecreateAPIRequest 淘宝短信签名创建接口 API请求
// taobao.jst.sms.signname.create
//
// 聚石塔短信签名创建接口
type TaobaojstsmssignnamecreateAPIRequest struct {
	model.Params
	// 创建签名入参
	_addSmsSignRequest *TopAddSmsSignRequest
}

// NewTaobaojstsmssignnamecreateRequest 初始化TaobaojstsmssignnamecreateAPIRequest对象
func NewTaobaojstsmssignnamecreateRequest() *TaobaojstsmssignnamecreateAPIRequest {
	return &TaobaojstsmssignnamecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmssignnamecreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.signname.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmssignnamecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmssignnamecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddSmsSignRequest is AddSmsSignRequest Setter
// 创建签名入参
func (r *TaobaojstsmssignnamecreateAPIRequest) SetAddSmsSignRequest(_addSmsSignRequest *TopAddSmsSignRequest) error {
	r._addSmsSignRequest = _addSmsSignRequest
	r.Set("add_sms_sign_request", _addSmsSignRequest)
	return nil
}

// GetAddSmsSignRequest AddSmsSignRequest Getter
func (r TaobaojstsmssignnamecreateAPIRequest) GetAddSmsSignRequest() *TopAddSmsSignRequest {
	return r._addSmsSignRequest
}
