package mc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mc"
)

/* 
订单信息同步 
tmall.mc.record.order.sync

订单信息同步(零售云接口)
*/
func TmallMcRecordOrderSync(clt *core.SDKClient, req *mc.TmallMcRecordOrderSyncAPIRequest, session string) (*mc.TmallMcRecordOrderSyncAPIResponse, error) {
    var resp mc.TmallMcRecordOrderSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
