package maitix

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
计算渠道用户下单快递费 API请求
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费
*/
type AlibabaDamaiMaitixDistributionDeliveryCalculateRequest struct {
    model.Params
    // 入参
    param   *OpenApiPostFeeParam
}

// 初始化AlibabaDamaiMaitixDistributionDeliveryCalculateRequest对象
func NewAlibabaDamaiMaitixDistributionDeliveryCalculateRequest() *AlibabaDamaiMaitixDistributionDeliveryCalculateRequest{
    return &AlibabaDamaiMaitixDistributionDeliveryCalculateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetApiMethodName() string {
    return "alibaba.damai.maitix.distribution.delivery.calculate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) SetParam(param *OpenApiPostFeeParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaDamaiMaitixDistributionDeliveryCalculateRequest) GetParam() *OpenApiPostFeeParam {
    return r.param
}
