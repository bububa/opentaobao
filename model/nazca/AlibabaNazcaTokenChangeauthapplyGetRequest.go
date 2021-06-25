package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取变更认证申请信息 APIRequest
alibaba.nazca.token.changeauthapply.get

根据token获取变更认证申请信息
*/
type AlibabaNazcaTokenChangeauthapplyGetRequest struct {
    model.Params

    // token
    token   string 

}

func NewAlibabaNazcaTokenChangeauthapplyGetRequest() *AlibabaNazcaTokenChangeauthapplyGetRequest{
    return &AlibabaNazcaTokenChangeauthapplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.changeauthapply.get"
}

func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNazcaTokenChangeauthapplyGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaNazcaTokenChangeauthapplyGetRequest) GetToken() string {
    return r.token
}

