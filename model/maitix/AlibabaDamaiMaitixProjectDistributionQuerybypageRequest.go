package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目分页查询项目列表服务 APIRequest
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务
*/
type AlibabaDamaiMaitixProjectDistributionQuerybypageRequest struct {
    model.Params

    // 入参param
    param   *ProjectPageParam 

}

func NewAlibabaDamaiMaitixProjectDistributionQuerybypageRequest() *AlibabaDamaiMaitixProjectDistributionQuerybypageRequest{
    return &AlibabaDamaiMaitixProjectDistributionQuerybypageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.project.distribution.querybypage"
}

func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) SetParam(param *ProjectPageParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixProjectDistributionQuerybypageRequest) GetParam() *ProjectPageParam {
    return r.param
}

