package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryProjectStatusByProjectId API请求
alibaba.damai.maitix.opengateway.project.status.query

queryProjectStatusByProjectId
*/
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest struct {
    model.Params
    // 入参dto
    disProjectStatusQueryParam   *DisProjectStatusQueryDto
}

// 初始化AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest对象
func NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.project.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisProjectStatusQueryParam Setter
// 入参dto
func (r *AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) SetDisProjectStatusQueryParam(disProjectStatusQueryParam *DisProjectStatusQueryDto) error {
    r.disProjectStatusQueryParam = disProjectStatusQueryParam
    r.Set("dis_project_status_query_param", disProjectStatusQueryParam)
    return nil
}

// DisProjectStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest) GetDisProjectStatusQueryParam() *DisProjectStatusQueryDto {
    return r.disProjectStatusQueryParam
}
