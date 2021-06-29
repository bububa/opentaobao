package xhotelonlineorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* 
酒店信用住订单状态更新 
taobao.xhotel.order.alipayface.update

完成对信用住或者面付订单的状态的更新。包含订单状态的确认，入离店状态的更新等等。（不适用于预付订单）
*/
func TaobaoXhotelOrderAlipayfaceUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceUpdateRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderAlipayfaceUpdateAPIResponse, error) {
    var resp xhotelonlineorder.TaobaoXhotelOrderAlipayfaceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
