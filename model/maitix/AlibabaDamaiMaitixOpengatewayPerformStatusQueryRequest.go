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
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest struct {
    model.Params
    // 入参
    _disPerformStatusQueryParam   *DisPerformStatusQueryDTO
}

// 初始化AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest对象
func NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest{
    return &AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.perform.status.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisPerformStatusQueryParam Setter
// 入参
func (r *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) SetDisPerformStatusQueryParam(_disPerformStatusQueryParam *DisPerformStatusQueryDTO) error {
    r._disPerformStatusQueryParam = _disPerformStatusQueryParam
    r.Set("dis_perform_status_query_param", _disPerformStatusQueryParam)
    return nil
}

// DisPerformStatusQueryParam Getter
func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIRequest) GetDisPerformStatusQueryParam() *DisPerformStatusQueryDTO {
    return r._disPerformStatusQueryParam
}
