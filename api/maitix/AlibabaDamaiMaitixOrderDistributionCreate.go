package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
大麦-新分销下单 
alibaba.damai.maitix.order.distribution.create

createDistributionOrder
*/
func AlibabaDamaiMaitixOrderDistributionCreate(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderDistributionCreateAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOrderDistributionCreateAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOrderDistributionCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
