package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦-新分销下单 APIRequest
alibaba.damai.maitix.order.distribution.create

createDistributionOrder
*/
type AlibabaDamaiMaitixOrderDistributionCreateRequest struct {
    model.Params

    // 下单参数param
    param   *MoaOrderParam 

}

func NewAlibabaDamaiMaitixOrderDistributionCreateRequest() *AlibabaDamaiMaitixOrderDistributionCreateRequest{
    return &AlibabaDamaiMaitixOrderDistributionCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.order.distribution.create"
}

func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixOrderDistributionCreateRequest) SetParam(param *MoaOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixOrderDistributionCreateRequest) GetParam() *MoaOrderParam {
    return r.param
}

