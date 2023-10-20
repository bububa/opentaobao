package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopauthtokenrefreshAPIRequest 刷新Access Token API请求
// taobao.top.auth.token.refresh
//
// 根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。
type TaobaotopauthtokenrefreshAPIRequest struct {
	model.Params
	// grantType==refresh_token 时需要
	_refreshToken string
}

// NewTaobaotopauthtokenrefreshRequest 初始化TaobaotopauthtokenrefreshAPIRequest对象
func NewTaobaotopauthtokenrefreshRequest() *TaobaotopauthtokenrefreshAPIRequest {
	return &TaobaotopauthtokenrefreshAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopauthtokenrefreshAPIRequest) GetApiMethodName() string {
	return "taobao.top.auth.token.refresh"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopauthtokenrefreshAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopauthtokenrefreshAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefreshToken is RefreshToken Setter
// grantType==refresh_token 时需要
func (r *TaobaotopauthtokenrefreshAPIRequest) SetRefreshToken(_refreshToken string) error {
	r._refreshToken = _refreshToken
	r.Set("refresh_token", _refreshToken)
	return nil
}

// GetRefreshToken RefreshToken Getter
func (r TaobaotopauthtokenrefreshAPIRequest) GetRefreshToken() string {
	return r._refreshToken
}
