package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销状态查询接口queryProjectStatusByProjectId 
alibaba.damai.maitix.opengateway.project.status.query

queryProjectStatusByProjectId
*/
func AlibabaDamaiMaitixOpengatewayProjectStatusQuery(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest, session string) (*maitix.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
