package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销单个项目信息查询 APIRequest
alibaba.damai.maitix.project.distribution.query

发布分销项目查询单个项目信息接口
*/
type AlibabaDamaiMaitixProjectDistributionQueryRequest struct {
    model.Params

    // 项目id
    projectId   int64 

}

func NewAlibabaDamaiMaitixProjectDistributionQueryRequest() *AlibabaDamaiMaitixProjectDistributionQueryRequest{
    return &AlibabaDamaiMaitixProjectDistributionQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixProjectDistributionQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.query"
}

func (r AlibabaDamaiMaitixProjectDistributionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixProjectDistributionQueryRequest) SetProjectId(projectId int64) error {
    r.projectId = projectId
    r.Set("project_id", projectId)
    return nil
}

func (r AlibabaDamaiMaitixProjectDistributionQueryRequest) GetProjectId() int64 {
    return r.projectId
}

