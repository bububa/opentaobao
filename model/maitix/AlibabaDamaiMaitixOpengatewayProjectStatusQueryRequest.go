package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryProjectStatusByProjectId APIRequest
alibaba.damai.maitix.opengateway.project.status.query

queryProjectStatusByProjectId
*/
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest struct {
    model.Params

    // 入参dto
    disProjectStatusQueryParam   *DisProjectStatusQueryDto 

}

func NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.project.status.query"
}

func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) SetDisProjectStatusQueryParam(disProjectStatusQueryParam *DisProjectStatusQueryDto) error {
    r.disProjectStatusQueryParam = disProjectStatusQueryParam
    r.Set("dis_project_status_query_param", disProjectStatusQueryParam)
    return nil
}

func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetDisProjectStatusQueryParam() *DisProjectStatusQueryDto {
    return r.disProjectStatusQueryParam
}

