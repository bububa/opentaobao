package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取开放平台出口IP段 APIRequest
taobao.top.ipout.get

获取开放平台出口IP段
*/
type TaobaoTopIpoutGetRequest struct {
    model.Params

}

func NewTaobaoTopIpoutGetRequest() *TaobaoTopIpoutGetRequest{
    return &TaobaoTopIpoutGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTopIpoutGetRequest) GetApiMethodName() string {
    return "taobao.top.ipout.get"
}

func (r TaobaoTopIpoutGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


