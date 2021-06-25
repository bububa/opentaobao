package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按用户获取支付宝代扣协议链接地址 APIRequest
taobao.oc.ap.contracturl.get

按用户获取支付宝代扣协议链接地址
*/
type TaobaoOcApContracturlGetRequest struct {
    model.Params

}

func NewTaobaoOcApContracturlGetRequest() *TaobaoOcApContracturlGetRequest{
    return &TaobaoOcApContracturlGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcApContracturlGetRequest) GetApiMethodName() string {
    return "taobao.oc.ap.contracturl.get"
}

func (r TaobaoOcApContracturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


