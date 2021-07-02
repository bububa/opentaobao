package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopAuthTokenRefreshAPIRequest 刷新Access Token API请求
// taobao.top.auth.token.refresh
//
// 根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。
type TaobaoTopAuthTokenRefreshAPIRequest struct {
	model.Params
	// grantType==refresh_token 时需要
	_refreshToken string
}

// NewTaobaoTopAuthTokenRefreshRequest 初始化TaobaoTopAuthTokenRefreshAPIRequest对象
func NewTaobaoTopAuthTokenRefreshRequest() *TaobaoTopAuthTokenRefreshAPIRequest {
	return &TaobaoTopAuthTokenRefreshAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetApiMethodName() string {
	return "taobao.top.auth.token.refresh"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefreshToken Setter
// grantType==refresh_token 时需要
func (r *TaobaoTopAuthTokenRefreshAPIRequest) SetRefreshToken(_refreshToken string) error {
	r._refreshToken = _refreshToken
	r.Set("refresh_token", _refreshToken)
	return nil
}

// Get RefreshToken Getter
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetRefreshToken() string {
	return r._refreshToken
}
