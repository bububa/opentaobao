package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
openmall订单支付 
taobao.openmall.trade.agreepay

openmall订单支付
*/
func TaobaoOpenmallTradeAgreepay(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeAgreepayRequest, session string) (*openmall.TaobaoOpenmallTradeAgreepayAPIResponse, error) {
    var resp openmall.TaobaoOpenmallTradeAgreepayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
