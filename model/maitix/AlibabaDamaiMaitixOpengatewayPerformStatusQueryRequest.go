package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryPerformStatusByPerformId APIRequest
alibaba.damai.maitix.opengateway.perform.status.query

queryPerformStatusByPerformId
*/
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest struct {
    model.Params

    // 入参
    disPerformStatusQueryParam   *DisPerformStatusQueryDto 

}

func NewAlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest{
    return &AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.opengateway.perform.status.query"
}

func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) SetDisPerformStatusQueryParam(disPerformStatusQueryParam *DisPerformStatusQueryDto) error {
    r.disPerformStatusQueryParam = disPerformStatusQueryParam
    r.Set("dis_perform_status_query_param", disPerformStatusQueryParam)
    return nil
}

func (r AlibabaDamaiMaitixOpengatewayPerformStatusQueryRequest) GetDisPerformStatusQueryParam() *DisPerformStatusQueryDto {
    return r.disPerformStatusQueryParam
}

