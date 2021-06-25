package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
入库通知单 
taobao.wlb.wms.stock.in.order.notify

入库通知单
*/
func TaobaoWlbWmsStockInOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockInOrderNotifyRequest, session string) (*wms.TaobaoWlbWmsStockInOrderNotifyResponse, error) {
    var resp wms.TaobaoWlbWmsStockInOrderNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
