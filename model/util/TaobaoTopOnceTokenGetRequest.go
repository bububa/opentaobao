package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
网关一次性token获取 APIRequest
taobao.top.once.token.get

网关一次性token获取
*/
type TaobaoTopOnceTokenGetRequest struct {
    model.Params

    // sec_token
    secToken   string 

}

func NewTaobaoTopOnceTokenGetRequest() *TaobaoTopOnceTokenGetRequest{
    return &TaobaoTopOnceTokenGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopOnceTokenGetRequest) GetApiMethodName() string {
    return "taobao.top.once.token.get"
}

func (r TaobaoTopOnceTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTopOnceTokenGetRequest) SetSecToken(secToken string) error {
    r.secToken = secToken
    r.Set("sec_token", secToken)
    return nil
}

func (r TaobaoTopOnceTokenGetRequest) GetSecToken() string {
    return r.secToken
}

