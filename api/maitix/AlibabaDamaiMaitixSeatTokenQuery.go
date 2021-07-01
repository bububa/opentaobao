package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销商选座获取qtoken 
alibaba.damai.maitix.seat.token.query

选座分销，分销商查询token
*/
func AlibabaDamaiMaitixSeatTokenQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixSeatTokenQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixSeatTokenQueryAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixSeatTokenQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
