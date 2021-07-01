package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
出库单通知 
taobao.wlb.wms.stock.out.order.notify

出库单通知
*/
func TaobaoWlbWmsStockOutOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockOutOrderNotifyAPIRequest, session string) (*wms.TaobaoWlbWmsStockOutOrderNotifyAPIResponse, error) {
    var resp wms.TaobaoWlbWmsStockOutOrderNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
