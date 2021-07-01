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
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest struct {
    model.Params
    // 入参dto
    _disProjectStatusQueryParam   *DisProjectStatusQueryDto
}

// 初始化AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayProjectStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest{
    return &AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.project.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisProjectStatusQueryParam Setter
// 入参dto
func (r *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) SetDisProjectStatusQueryParam(_disProjectStatusQueryParam *DisProjectStatusQueryDto) error {
    r._disProjectStatusQueryParam = _disProjectStatusQueryParam
    r.Set("dis_project_status_query_param", _disProjectStatusQueryParam)
    return nil
}

// DisProjectStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIRequest) GetDisProjectStatusQueryParam() *DisProjectStatusQueryDto {
    return r._disProjectStatusQueryParam
}
