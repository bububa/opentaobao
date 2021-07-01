package rhino

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/rhino"
)

/* 
WMS003成衣入库确认 
taobao.rhino.supplychain.inbound.confirm

【WMS003】【同步成衣入库完成信息】
*/
func TaobaoRhinoSupplychainInboundConfirm(clt *core.SDKClient, req *rhino.TaobaoRhinoSupplychainInboundConfirmAPIRequest, session string) (*rhino.TaobaoRhinoSupplychainInboundConfirmAPIResponse, error) {
    var resp rhino.TaobaoRhinoSupplychainInboundConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
