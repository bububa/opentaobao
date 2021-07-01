package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销单个项目信息查询 API请求
alibaba.damai.maitix.project.distribution.query

发布分销项目查询单个项目信息接口
*/
type AlibabaDamaiMaitixProjectDistributionQueryAPIRequest struct {
    model.Params
    // 项目id
    _projectId   int64
}

// 初始化AlibabaDamaiMaitixProjectDistributionQueryAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQueryRequest() *AlibabaDamaiMaitixProjectDistributionQueryAPIRequest{
    return &AlibabaDamaiMaitixProjectDistributionQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectId Setter
// 项目id
func (r *AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) SetProjectId(_projectId int64) error {
    r._projectId = _projectId
    r.Set("project_id", _projectId)
    return nil
}

// ProjectId Getter
func (r AlibabaDamaiMaitixProjectDistributionQueryAPIRequest) GetProjectId() int64 {
    return r._projectId
}
