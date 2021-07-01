package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
大麦-直接退票 
alibaba.damai.maitix.order.directrefund

大麦-退票
*/
func AlibabaDamaiMaitixOrderDirectrefund(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderDirectrefundAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOrderDirectrefundAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOrderDirectrefundAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
