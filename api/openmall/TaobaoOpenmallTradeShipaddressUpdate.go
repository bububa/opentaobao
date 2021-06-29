package openmall

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openmall"
)

/* 
Openmall订单收货地址修改 
taobao.openmall.trade.shipaddress.update

Openmall订单收货地址修改
*/
func TaobaoOpenmallTradeShipaddressUpdate(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeShipaddressUpdateRequest, session string) (*openmall.TaobaoOpenmallTradeShipaddressUpdateAPIResponse, error) {
    var resp openmall.TaobaoOpenmallTradeShipaddressUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
