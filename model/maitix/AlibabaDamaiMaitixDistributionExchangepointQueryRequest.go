package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销查询取票点接口 API请求
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口
*/
type AlibabaDamaiMaitixDistributionExchangepointQueryRequest struct {
    model.Params
    // 必填-项目id
    _projectId   int64
}

// 初始化AlibabaDamaiMaitixDistributionExchangepointQueryRequest对象
func NewAlibabaDamaiMaitixDistributionExchangepointQueryRequest() *AlibabaDamaiMaitixDistributionExchangepointQueryRequest{
    return &AlibabaDamaiMaitixDistributionExchangepointQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.exchangepoint.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectId Setter
// 必填-项目id
func (r *AlibabaDamaiMaitixDistributionExchangepointQueryRequest) SetProjectId(_projectId int64) error {
    r._projectId = _projectId
    r.Set("project_id", _projectId)
    return nil
}

// ProjectId Getter
func (r AlibabaDamaiMaitixDistributionExchangepointQueryRequest) GetProjectId() int64 {
    return r._projectId
}
