package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销查询取票点接口 APIRequest
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口
*/
type AlibabaDamaiMaitixDistributionExchangepointQueryRequest struct {
    model.Params

    // 必填-项目id
    projectId   int64 

}

func NewAlibabaDamaiMaitixDistributionExchangepointQueryRequest() *AlibabaDamaiMaitixDistributionExchangepointQueryRequest{
    return &AlibabaDamaiMaitixDistributionExchangepointQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.exchangepoint.query"
}

func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixDistributionExchangepointQueryRequest) SetProjectId(projectId int64) error {
    r.projectId = projectId
    r.Set("project_id", projectId)
    return nil
}

func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetProjectId() int64 {
    return r.projectId
}

