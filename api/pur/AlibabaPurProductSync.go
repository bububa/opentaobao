package pur

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/pur"
)

/* 
同步产品 
alibaba.pur.product.sync

同步产品
*/
func AlibabaPurProductSync(clt *core.SDKClient, req *pur.AlibabaPurProductSyncRequest, session string) (*pur.AlibabaPurProductSyncAPIResponse, error) {
    var resp pur.AlibabaPurProductSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
