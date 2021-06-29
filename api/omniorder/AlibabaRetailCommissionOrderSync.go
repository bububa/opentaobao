package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
分佣数据传输 
alibaba.retail.commission.order.sync

同步分佣结果
*/
func AlibabaRetailCommissionOrderSync(clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionOrderSyncRequest, session string) (*omniorder.AlibabaRetailCommissionOrderSyncAPIResponse, error) {
    var resp omniorder.AlibabaRetailCommissionOrderSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
