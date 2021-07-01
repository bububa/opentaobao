package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
创建订单 
taobao.openmall.trade.create

创建Openmall订单
*/
func TaobaoOpenmallTradeCreate(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeCreateAPIRequest, session string) (*openmall.TaobaoOpenmallTradeCreateAPIResponse, error) {
    var resp openmall.TaobaoOpenmallTradeCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
