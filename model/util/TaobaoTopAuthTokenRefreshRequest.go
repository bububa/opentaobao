package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
刷新Access Token API请求
taobao.top.auth.token.refresh

根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。
*/
type TaobaoTopAuthTokenRefreshRequest struct {
    model.Params
    // grantType==refresh_token 时需要
    refreshToken   string
}

// 初始化TaobaoTopAuthTokenRefreshRequest对象
func NewTaobaoTopAuthTokenRefreshRequest() *TaobaoTopAuthTokenRefreshRequest{
    return &TaobaoTopAuthTokenRefreshRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTopAuthTokenRefreshRequest) GetApiMethodName() string {
    return "taobao.top.auth.token.refresh"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTopAuthTokenRefreshRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefreshToken Setter
// grantType==refresh_token 时需要
func (r *TaobaoTopAuthTokenRefreshRequest) SetRefreshToken(refreshToken string) error {
    r.refreshToken = refreshToken
    r.Set("refresh_token", refreshToken)
    return nil
}

// RefreshToken Getter
func (r TaobaoTopAuthTokenRefreshRequest) GetRefreshToken() string {
    return r.refreshToken
}
