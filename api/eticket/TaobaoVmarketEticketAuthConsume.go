package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
核销放行的核销接口 
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
*/
func TaobaoVmarketEticketAuthConsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketAuthConsumeAPIRequest, session string) (*eticket.TaobaoVmarketEticketAuthConsumeAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketAuthConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
