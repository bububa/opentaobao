package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
商家电子凭证发码成功回调接口 
taobao.vmarket.eticket.send

外部商家成功发码回调接口
*/
func TaobaoVmarketEticketSend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketSendRequest, session string) (*eticket.TaobaoVmarketEticketSendAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
