package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除奇门订单链路用户 APIRequest
taobao.qimen.trade.user.delete

删除奇门订单链路用户
*/
type TaobaoQimenTradeUserDeleteRequest struct {
    model.Params

}

func NewTaobaoQimenTradeUserDeleteRequest() *TaobaoQimenTradeUserDeleteRequest{
    return &TaobaoQimenTradeUserDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTradeUserDeleteRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.user.delete"
}

func (r TaobaoQimenTradeUserDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


