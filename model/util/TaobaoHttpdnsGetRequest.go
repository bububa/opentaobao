package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TOPDNS配置 APIRequest
taobao.httpdns.get

获取TOP DNS配置
*/
type TaobaoHttpdnsGetRequest struct {
    model.Params

}

func NewTaobaoHttpdnsGetRequest() *TaobaoHttpdnsGetRequest{
    return &TaobaoHttpdnsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoHttpdnsGetRequest) GetApiMethodName() string {
    return "taobao.httpdns.get"
}

func (r TaobaoHttpdnsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


