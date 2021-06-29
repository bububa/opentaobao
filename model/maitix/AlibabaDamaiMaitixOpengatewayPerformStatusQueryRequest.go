package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryPerformStatusByPerformId API请求
alibaba.damai.maitix.opengateway.perform.status.query

queryPerformStatusByPerformId
*/
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest struct {
    model.Params
    // 入参
    disPerformStatusQueryParam   *DisPerformStatusQueryDto
}

// 初始化AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest对象
func NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.perform.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisPerformStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) SetDisPerformStatusQueryParam(disPerformStatusQueryParam *DisPerformStatusQueryDto) error {
    r.disPerformStatusQueryParam = disPerformStatusQueryParam
    r.Set("dis_perform_status_query_param", disPerformStatusQueryParam)
    return nil
}

// DisPerformStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetDisPerformStatusQueryParam() *DisPerformStatusQueryDto {
    return r.disPerformStatusQueryParam
}
