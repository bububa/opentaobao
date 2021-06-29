package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算渠道用户下单快递费 APIRequest
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费
*/
type AlibabaDamaiMaitixDistributionDeliveryCalculateRequest struct {
    model.Params

    // 入参
    param   *OpenApiPostFeeParam 

}

func NewAlibabaDamaiMaitixDistributionDeliveryCalculateRequest() *AlibabaDamaiMaitixDistributionDeliveryCalculateRequest{
    return &AlibabaDamaiMaitixDistributionDeliveryCalculateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.delivery.calculate"
}

func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) SetParam(param *OpenApiPostFeeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetParam() *OpenApiPostFeeParam {
    return r.param
}

