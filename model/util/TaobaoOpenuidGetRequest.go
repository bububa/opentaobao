package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权账号对应的OpenUid APIRequest
taobao.openuid.get

获取授权账号对应的OpenUid
*/
type TaobaoOpenuidGetRequest struct {
    model.Params

}

func NewTaobaoOpenuidGetRequest() *TaobaoOpenuidGetRequest{
    return &TaobaoOpenuidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenuidGetRequest) GetApiMethodName() string {
    return "taobao.openuid.get"
}

func (r TaobaoOpenuidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


