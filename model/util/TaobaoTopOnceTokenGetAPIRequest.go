package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOnceTokenGetAPIRequest 网关一次性token获取 API请求
// taobao.top.once.token.get
//
// 网关一次性token获取
type TaobaoTopOnceTokenGetAPIRequest struct {
	model.Params
	// sec_token
	_secToken string
}

// NewTaobaoTopOnceTokenGetRequest 初始化TaobaoTopOnceTokenGetAPIRequest对象
func NewTaobaoTopOnceTokenGetRequest() *TaobaoTopOnceTokenGetAPIRequest {
	return &TaobaoTopOnceTokenGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOnceTokenGetAPIRequest) GetApiMethodName() string {
	return "taobao.top.once.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOnceTokenGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SecToken Setter
// sec_token
func (r *TaobaoTopOnceTokenGetAPIRequest) SetSecToken(_secToken string) error {
	r._secToken = _secToken
	r.Set("sec_token", _secToken)
	return nil
}

// Get SecToken Getter
func (r TaobaoTopOnceTokenGetAPIRequest) GetSecToken() string {
	return r._secToken
}
