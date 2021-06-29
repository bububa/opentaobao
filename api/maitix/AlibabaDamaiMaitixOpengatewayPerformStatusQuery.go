package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销状态查询接口queryPerformStatusByPerformId 
alibaba.damai.maitix.opengateway.perform.status.query

queryPerformStatusByPerformId
*/
func AlibabaDamaiMaitixOpengatewayPerformStatusQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest, session string) (*maitix.AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
