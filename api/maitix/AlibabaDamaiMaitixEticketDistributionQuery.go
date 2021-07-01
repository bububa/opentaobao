package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销电子票查询接口 
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口
*/
func AlibabaDamaiMaitixEticketDistributionQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixEticketDistributionQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixEticketDistributionQueryAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixEticketDistributionQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
