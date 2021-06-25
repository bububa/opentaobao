package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据获取压测用户的sessionKey APIRequest
taobao.streetest.session.get

根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
*/
type TaobaoStreetestSessionGetRequest struct {
    model.Params

}

func NewTaobaoStreetestSessionGetRequest() *TaobaoStreetestSessionGetRequest{
    return &TaobaoStreetestSessionGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoStreetestSessionGetRequest) GetApiMethodName() string {
    return "taobao.streetest.session.get"
}

func (r TaobaoStreetestSessionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


