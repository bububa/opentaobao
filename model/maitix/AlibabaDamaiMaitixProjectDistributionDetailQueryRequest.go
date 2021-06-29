package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦分销项目内容详情查询 APIRequest
alibaba.damai.maitix.project.distribution.detail.query

大麦分销项目内容详情查询，提供项目的内容详情
*/
type AlibabaDamaiMaitixProjectDistributionDetailQueryRequest struct {
    model.Params

    // 项目ID，前提已授权
    projectId   int64 

}

func NewAlibabaDamaiMaitixProjectDistributionDetailQueryRequest() *AlibabaDamaiMaitixProjectDistributionDetailQueryRequest{
    return &AlibabaDamaiMaitixProjectDistributionDetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixProjectDistributionDetailQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.detail.query"
}

func (r AlibabaDamaiMaitixProjectDistributionDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixProjectDistributionDetailQueryRequest) SetProjectId(projectId int64) error {
    r.projectId = projectId
    r.Set("project_id", projectId)
    return nil
}

func (r AlibabaDamaiMaitixProjectDistributionDetailQueryRequest) GetProjectId() int64 {
    return r.projectId
}

