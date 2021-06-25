package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取ISV发起请求服务器IP APIRequest
taobao.appip.get

获取ISV发起请求服务器IP
*/
type TaobaoAppipGetRequest struct {
    model.Params

}

func NewTaobaoAppipGetRequest() *TaobaoAppipGetRequest{
    return &TaobaoAppipGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppipGetRequest) GetApiMethodName() string {
    return "taobao.appip.get"
}

func (r TaobaoAppipGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


