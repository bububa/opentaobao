package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
销售退货通知 
taobao.wlb.wms.return.order.notify

销售退货通知
*/
func TaobaoWlbWmsReturnOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsReturnOrderNotifyRequest, session string) (*wms.TaobaoWlbWmsReturnOrderNotifyResponse, error) {
    var resp wms.TaobaoWlbWmsReturnOrderNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
