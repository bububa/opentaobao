package wms

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wms"
)

/* 
发货订单通知 
taobao.wlb.wms.consign.order.notify

发货订单通知
*/
func TaobaoWlbWmsConsignOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsConsignOrderNotifyRequest, session string) (*wms.TaobaoWlbWmsConsignOrderNotifyAPIResponse, error) {
    var resp wms.TaobaoWlbWmsConsignOrderNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
