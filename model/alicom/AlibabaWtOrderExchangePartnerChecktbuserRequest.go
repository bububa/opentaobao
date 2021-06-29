package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分兑换校验淘宝账号是否存在 API请求
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在
*/
type AlibabaWtOrderExchangePartnerChecktbuserRequest struct {
    model.Params
    // model入参
    outExchangeModel   *OutExchangeModel
}

// 初始化AlibabaWtOrderExchangePartnerChecktbuserRequest对象
func NewAlibabaWtOrderExchangePartnerChecktbuserRequest() *AlibabaWtOrderExchangePartnerChecktbuserRequest{
    return &AlibabaWtOrderExchangePartnerChecktbuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetApiMethodName() string {
    return "alibaba.wt.order.exchange.partner.checktbuser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutExchangeModel Setter
// model入参
func (r *AlibabaWtOrderExchangePartnerChecktbuserRequest) SetOutExchangeModel(outExchangeModel *OutExchangeModel) error {
    r.outExchangeModel = outExchangeModel
    r.Set("out_exchange_model", outExchangeModel)
    return nil
}

// OutExchangeModel Getter
func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetOutExchangeModel() *OutExchangeModel {
    return r.outExchangeModel
}
