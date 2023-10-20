package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotoponcetokengetAPIRequest 网关一次性token获取 API请求
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
type TaobaotoponcetokengetAPIRequest struct {
	model.Params
	// sec_token
	_secToken string
}

// NewTaobaotoponcetokengetRequest 初始化TaobaotoponcetokengetAPIRequest对象
func NewTaobaotoponcetokengetRequest() *TaobaotoponcetokengetAPIRequest {
	return &TaobaotoponcetokengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotoponcetokengetAPIRequest) GetApiMethodName() string {
	return "taobao.top.once.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotoponcetokengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotoponcetokengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecToken is SecToken Setter
// sec_token
func (r *TaobaotoponcetokengetAPIRequest) SetSecToken(_secToken string) error {
	r._secToken = _secToken
	r.Set("sec_token", _secToken)
	return nil
}

// GetSecToken SecToken Getter
func (r TaobaotoponcetokengetAPIRequest) GetSecToken() string {
	return r._secToken
}
