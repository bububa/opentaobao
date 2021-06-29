package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目分页查询项目列表服务 API请求
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务
*/
type AlibabaDamaiMaitixProjectDistributionQuerybypageRequest struct {
    model.Params
    // 入参param
    param   *ProjectPageParam
}

// 初始化AlibabaDamaiMaitixProjectDistributionQuerybypageRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerybypageRequest() *AlibabaDamaiMaitixProjectDistributionQuerybypageRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerybypageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querybypage"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参param
func (r *AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) SetParam(param *ProjectPageParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetParam() *ProjectPageParam {
    return r.param
}
