package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取变更认证申请信息 API请求
alibaba.nazca.token.changeauthapply.get

根据token获取变更认证申请信息
*/
type AlibabaNazcaTokenChangeauthapplyGetRequest struct {
    model.Params
    // token
    _token   string
}

// 初始化AlibabaNazcaTokenChangeauthapplyGetRequest对象
func NewAlibabaNazcaTokenChangeauthapplyGetRequest() *AlibabaNazcaTokenChangeauthapplyGetRequest{
    return &AlibabaNazcaTokenChangeauthapplyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.changeauthapply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// token
func (r *AlibabaNazcaTokenChangeauthapplyGetRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetToken() string {
    return r._token
}
