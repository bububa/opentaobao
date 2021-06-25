package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
是否为新支付订单 
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
*/
func AlibabaMosOnsiteTradeIsnewpayorder(clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradeIsnewpayorderRequest, session string) (*mos.AlibabaMosOnsiteTradeIsnewpayorderResponse, error) {
    var resp mos.AlibabaMosOnsiteTradeIsnewpayorderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
