package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccounttokenvalidateAPIRequest open account token验证 API请求
// taobao.open.account.token.validate
//
// open account token验证
type TaobaoopenaccounttokenvalidateAPIRequest struct {
	model.Params
	// token
	_paramToken string
}

// NewTaobaoopenaccounttokenvalidateRequest 初始化TaobaoopenaccounttokenvalidateAPIRequest对象
func NewTaobaoopenaccounttokenvalidateRequest() *TaobaoopenaccounttokenvalidateAPIRequest {
	return &TaobaoopenaccounttokenvalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccounttokenvalidateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.token.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccounttokenvalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccounttokenvalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamToken is ParamToken Setter
// token
func (r *TaobaoopenaccounttokenvalidateAPIRequest) SetParamToken(_paramToken string) error {
	r._paramToken = _paramToken
	r.Set("param_token", _paramToken)
	return nil
}

// GetParamToken ParamToken Getter
func (r TaobaoopenaccounttokenvalidateAPIRequest) GetParamToken() string {
	return r._paramToken
}
