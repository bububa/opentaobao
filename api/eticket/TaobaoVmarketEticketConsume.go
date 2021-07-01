package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子票券消费通知 
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口
*/
func TaobaoVmarketEticketConsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketConsumeAPIRequest, session string) (*eticket.TaobaoVmarketEticketConsumeAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
