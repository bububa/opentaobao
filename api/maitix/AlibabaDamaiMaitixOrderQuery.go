package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
大麦-查询分销单 
alibaba.damai.maitix.order.query

查询分销单
*/
func AlibabaDamaiMaitixOrderQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderQueryRequest, session string) (*maitix.AlibabaDamaiMaitixOrderQueryAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
