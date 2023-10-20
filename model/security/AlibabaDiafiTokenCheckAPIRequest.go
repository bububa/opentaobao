package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadiafitokencheckAPIRequest 天朗token校验API API请求
// alibaba.diafi.token.check
//
// 天朗token校验
type AlibabadiafitokencheckAPIRequest struct {
	model.Params
	// token
	_token string
	// 对应配置
	_app string
}

// NewAlibabadiafitokencheckRequest 初始化AlibabadiafitokencheckAPIRequest对象
func NewAlibabadiafitokencheckRequest() *AlibabadiafitokencheckAPIRequest {
	return &AlibabadiafitokencheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadiafitokencheckAPIRequest) GetApiMethodName() string {
	return "alibaba.diafi.token.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadiafitokencheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadiafitokencheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabadiafitokencheckAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabadiafitokencheckAPIRequest) GetToken() string {
	return r._token
}

// SetApp is App Setter
// 对应配置
func (r *AlibabadiafitokencheckAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// GetApp App Getter
func (r AlibabadiafitokencheckAPIRequest) GetApp() string {
	return r._app
}
