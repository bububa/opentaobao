package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
无法发码回调 
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
func TaobaoVmarketEticketFailsend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFailsendRequest, session string) (*eticket.TaobaoVmarketEticketFailsendAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketFailsendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
