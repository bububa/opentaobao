package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销项目分页查询项目列表服务 
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务
*/
func AlibabaDamaiMaitixProjectDistributionQuerybypage(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageRequest, session string) (*maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
