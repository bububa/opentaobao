package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDiafiTokenCheckAPIRequest 天朗token校验API API请求
// alibaba.diafi.token.check
//
// 天朗token校验
type AlibabaDiafiTokenCheckAPIRequest struct {
	model.Params
	// token
	_token string
	// 对应配置
	_app string
}

// NewAlibabaDiafiTokenCheckRequest 初始化AlibabaDiafiTokenCheckAPIRequest对象
func NewAlibabaDiafiTokenCheckRequest() *AlibabaDiafiTokenCheckAPIRequest {
	return &AlibabaDiafiTokenCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDiafiTokenCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.diafi.token.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDiafiTokenCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDiafiTokenCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// token
func (r *AlibabaDiafiTokenCheckAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaDiafiTokenCheckAPIRequest) GetToken() string {
	return r._token
}

// SetApp is App Setter
// 对应配置
func (r *AlibabaDiafiTokenCheckAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// GetApp App Getter
func (r AlibabaDiafiTokenCheckAPIRequest) GetApp() string {
	return r._app
}
