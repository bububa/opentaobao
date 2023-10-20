package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscrighttokencheckAPIRequest 实物奖品凭证校验 API请求
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
type AlibabaalscrighttokencheckAPIRequest struct {
	model.Params
	// 实物奖品凭证
	_token string
}

// NewAlibabaalscrighttokencheckRequest 初始化AlibabaalscrighttokencheckAPIRequest对象
func NewAlibabaalscrighttokencheckRequest() *AlibabaalscrighttokencheckAPIRequest {
	return &AlibabaalscrighttokencheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscrighttokencheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.right.token.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscrighttokencheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscrighttokencheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 实物奖品凭证
func (r *AlibabaalscrighttokencheckAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaalscrighttokencheckAPIRequest) GetToken() string {
	return r._token
}
