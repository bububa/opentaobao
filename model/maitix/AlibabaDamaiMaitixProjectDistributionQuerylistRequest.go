package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目列表查询（已过时，不推荐使用） API请求
alibaba.damai.maitix.project.distribution.querylist

分销项目列表查询接口（已过时，不推荐使用）
*/
type AlibabaDamaiMaitixProjectDistributionQuerylistRequest struct {
    model.Params
}

// 初始化AlibabaDamaiMaitixProjectDistributionQuerylistRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest() *AlibabaDamaiMaitixProjectDistributionQuerylistRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerylistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerylistRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querylist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerylistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
