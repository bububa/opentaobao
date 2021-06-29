package ship

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ship"
)

/* 
船票班次变更回调 
alitrip.ship.product.syncnunber

船票班次变更回调
*/
func AlitripShipProductSyncnunber(clt *core.SDKClient, req *ship.AlitripShipProductSyncnunberRequest, session string) (*ship.AlitripShipProductSyncnunberAPIResponse, error) {
    var resp ship.AlitripShipProductSyncnunberAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}