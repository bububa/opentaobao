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
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest struct {
    model.Params
}

// 初始化AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest() *AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querylist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
