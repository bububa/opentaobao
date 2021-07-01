package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
大麦-库存释放 
alibaba.damai.maitix.order.cancel

库存释放
*/
func AlibabaDamaiMaitixOrderCancel(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderCancelAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOrderCancelAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOrderCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
