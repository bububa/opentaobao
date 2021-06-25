package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取认证申请信息 APIRequest
alibaba.nazca.token.authapply.get

根据token获取认证申请信息
*/
type AlibabaNazcaTokenAuthapplyGetRequest struct {
    model.Params

    // token
    token   string 

}

func NewAlibabaNazcaTokenAuthapplyGetRequest() *AlibabaNazcaTokenAuthapplyGetRequest{
    return &AlibabaNazcaTokenAuthapplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNazcaTokenAuthapplyGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.authapply.get"
}

func (r AlibabaNazcaTokenAuthapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNazcaTokenAuthapplyGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaNazcaTokenAuthapplyGetRequest) GetToken() string {
    return r.token
}

