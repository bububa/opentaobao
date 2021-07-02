package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountTokenValidateAPIRequest open account token验证 API请求
// taobao.open.account.token.validate
//
// open account token验证
type TaobaoOpenAccountTokenValidateAPIRequest struct {
	model.Params
	// token
	_paramToken string
}

// NewTaobaoOpenAccountTokenValidateRequest 初始化TaobaoOpenAccountTokenValidateAPIRequest对象
func NewTaobaoOpenAccountTokenValidateRequest() *TaobaoOpenAccountTokenValidateAPIRequest {
	return &TaobaoOpenAccountTokenValidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.token.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamToken Setter
// token
func (r *TaobaoOpenAccountTokenValidateAPIRequest) SetParamToken(_paramToken string) error {
	r._paramToken = _paramToken
	r.Set("param_token", _paramToken)
	return nil
}

// Get ParamToken Getter
func (r TaobaoOpenAccountTokenValidateAPIRequest) GetParamToken() string {
	return r._paramToken
}
