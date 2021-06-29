package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目列表查询（已过时，不推荐使用） APIRequest
alibaba.damai.maitix.project.distribution.querylist

分销项目列表查询接口（已过时，不推荐使用）
*/
type AlibabaDamaiMaitixProjectDistributionQuerylistRequest struct {
    model.Params

}

func NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest() *AlibabaDamaiMaitixProjectDistributionQuerylistRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerylistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixProjectDistributionQuerylistRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querylist"
}

func (r AlibabaDamaiMaitixProjectDistributionQuerylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


