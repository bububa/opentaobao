package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得当前登录用户的授权账户列表 APIRequest
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
type TaobaoSimbaCustomersAuthorizedGetRequest struct {
    model.Params

}

func NewTaobaoSimbaCustomersAuthorizedGetRequest() *TaobaoSimbaCustomersAuthorizedGetRequest{
    return &TaobaoSimbaCustomersAuthorizedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCustomersAuthorizedGetRequest) GetApiMethodName() string {
    return "taobao.simba.customers.authorized.get"
}

func (r TaobaoSimbaCustomersAuthorizedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


