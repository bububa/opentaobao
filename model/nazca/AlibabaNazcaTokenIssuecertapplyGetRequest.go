package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据token获取出证申请信息 APIRequest
alibaba.nazca.token.issuecertapply.get

根据token获取出证申请信息
*/
type AlibabaNazcaTokenIssuecertapplyGetRequest struct {
    model.Params

    // token
    token   string 

}

func NewAlibabaNazcaTokenIssuecertapplyGetRequest() *AlibabaNazcaTokenIssuecertapplyGetRequest{
    return &AlibabaNazcaTokenIssuecertapplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.issuecertapply.get"
}

func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNazcaTokenIssuecertapplyGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaNazcaTokenIssuecertapplyGetRequest) GetToken() string {
    return r.token
}

