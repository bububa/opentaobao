package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
大麦-出票 
alibaba.damai.maitix.order.confirm

出票
*/
func AlibabaDamaiMaitixOrderConfirm(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderConfirmAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOrderConfirmAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOrderConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
