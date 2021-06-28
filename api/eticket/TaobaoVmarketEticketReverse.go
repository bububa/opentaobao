package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
电子凭证冲正接口 
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口
*/
func TaobaoVmarketEticketReverse(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketReverseRequest, session string) (*eticket.TaobaoVmarketEticketReverseAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketReverseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
