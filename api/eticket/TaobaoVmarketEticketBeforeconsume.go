package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子凭证验码前置确认 
taobao.vmarket.eticket.beforeconsume

商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
*/
func TaobaoVmarketEticketBeforeconsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketBeforeconsumeAPIRequest, session string) (*eticket.TaobaoVmarketEticketBeforeconsumeAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketBeforeconsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
