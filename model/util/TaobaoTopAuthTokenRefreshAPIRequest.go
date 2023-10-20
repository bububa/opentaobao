package util

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopAuthTokenRefreshAPIRequest) Reset() {
	r._refreshToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetApiMethodName() string {
	return "taobao.top.auth.token.refresh"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefreshToken is RefreshToken Setter
// grantType==refresh_token 时需要
func (r *TaobaoTopAuthTokenRefreshAPIRequest) SetRefreshToken(_refreshToken string) error {
	r._refreshToken = _refreshToken
	r.Set("refresh_token", _refreshToken)
	return nil
}

// GetRefreshToken RefreshToken Getter
func (r TaobaoTopAuthTokenRefreshAPIRequest) GetRefreshToken() string {
	return r._refreshToken
}

var poolTaobaoTopAuthTokenRefreshAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopAuthTokenRefreshRequest()
	},
}

// GetTaobaoTopAuthTokenRefreshRequest 从 sync.Pool 获取 TaobaoTopAuthTokenRefreshAPIRequest
func GetTaobaoTopAuthTokenRefreshAPIRequest() *TaobaoTopAuthTokenRefreshAPIRequest {
	return poolTaobaoTopAuthTokenRefreshAPIRequest.Get().(*TaobaoTopAuthTokenRefreshAPIRequest)
}

// ReleaseTaobaoTopAuthTokenRefreshAPIRequest 将 TaobaoTopAuthTokenRefreshAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopAuthTokenRefreshAPIRequest(v *TaobaoTopAuthTokenRefreshAPIRequest) {
	v.Reset()
	poolTaobaoTopAuthTokenRefreshAPIRequest.Put(v)
}
