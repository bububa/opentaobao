package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息获取 APIRequest
taobao.jds.hluser.get

订单全链路用户信息获取
*/
type TaobaoJdsHluserGetRequest struct {
    model.Params

}

func NewTaobaoJdsHluserGetRequest() *TaobaoJdsHluserGetRequest{
    return &TaobaoJdsHluserGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJdsHluserGetRequest) GetApiMethodName() string {
    return "taobao.jds.hluser.get"
}

func (r TaobaoJdsHluserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


