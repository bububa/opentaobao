package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取淘宝系统当前时间 APIRequest
taobao.time.get

获取淘宝系统当前时间
*/
type TaobaoTimeGetRequest struct {
    model.Params

}

func NewTaobaoTimeGetRequest() *TaobaoTimeGetRequest{
    return &TaobaoTimeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTimeGetRequest) GetApiMethodName() string {
    return "taobao.time.get"
}

func (r TaobaoTimeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


