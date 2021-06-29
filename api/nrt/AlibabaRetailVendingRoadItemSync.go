package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
贩卖机库存商品同步 
alibaba.retail.vending.road.item.sync

贩卖机库存商品同步
*/
func AlibabaRetailVendingRoadItemSync(clt *core.SDKClient, req *nrt.AlibabaRetailVendingRoadItemSyncRequest, session string) (*nrt.AlibabaRetailVendingRoadItemSyncAPIResponse, error) {
    var resp nrt.AlibabaRetailVendingRoadItemSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
