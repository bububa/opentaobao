package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenlinksessiongetAPIRequest 获取授权session信息 API请求
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
type TaobaoopenlinksessiongetAPIRequest struct {
	model.Params
	// 授权码
	_code string
}

// NewTaobaoopenlinksessiongetRequest 初始化TaobaoopenlinksessiongetAPIRequest对象
func NewTaobaoopenlinksessiongetRequest() *TaobaoopenlinksessiongetAPIRequest {
	return &TaobaoopenlinksessiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenlinksessiongetAPIRequest) GetApiMethodName() string {
	return "taobao.openlink.session.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenlinksessiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenlinksessiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 授权码
func (r *TaobaoopenlinksessiongetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoopenlinksessiongetAPIRequest) GetCode() string {
	return r._code
}
