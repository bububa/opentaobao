package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单获取对应买家的openUID APIRequest
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权
*/
type TaobaoOpenuidGetBytradeRequest struct {
    model.Params

    // 订单ID
    tid   int64 

}

func NewTaobaoOpenuidGetBytradeRequest() *TaobaoOpenuidGetBytradeRequest{
    return &TaobaoOpenuidGetBytradeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenuidGetBytradeRequest) GetApiMethodName() string {
    return "taobao.openuid.get.bytrade"
}

func (r TaobaoOpenuidGetBytradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenuidGetBytradeRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoOpenuidGetBytradeRequest) GetTid() int64 {
    return r.tid
}

