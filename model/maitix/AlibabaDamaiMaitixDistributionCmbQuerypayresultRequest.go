package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询招行支付状态api APIRequest
alibaba.damai.maitix.distribution.cmb.querypayresult

queryPayResult
*/
type AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest struct {
    model.Params

    // 入参param
    param   *QueryPayResultParam 

}

func NewAlibabaDamaiMaitixDistributionCmbQuerypayresultRequest() *AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest{
    return &AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.cmb.querypayresult"
}

func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) SetParam(param *QueryPayResultParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixDistributionCmbQuerypayresultRequest) GetParam() *QueryPayResultParam {
    return r.param
}

