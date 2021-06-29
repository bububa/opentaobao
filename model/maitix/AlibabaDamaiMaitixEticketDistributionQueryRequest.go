package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销电子票查询接口 APIRequest
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口
*/
type AlibabaDamaiMaitixEticketDistributionQueryRequest struct {
    model.Params

    // 入参param
    param   *EticketQueryParam 

}

func NewAlibabaDamaiMaitixEticketDistributionQueryRequest() *AlibabaDamaiMaitixEticketDistributionQueryRequest{
    return &AlibabaDamaiMaitixEticketDistributionQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.eticket.distribution.query"
}

func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixEticketDistributionQueryRequest) SetParam(param *EticketQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixEticketDistributionQueryRequest) GetParam() *EticketQueryParam {
    return r.param
}

