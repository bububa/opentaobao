package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
分销项目列表查询（已过时，不推荐使用） 
alibaba.damai.maitix.project.distribution.querylist

分销项目列表查询接口（已过时，不推荐使用）
*/
func AlibabaDamaiMaitixProjectDistributionQuerylist(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQuerylistRequest, session string) (*maitix.AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
