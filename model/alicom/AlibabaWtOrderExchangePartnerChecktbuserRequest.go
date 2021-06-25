package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
积分兑换校验淘宝账号是否存在 APIRequest
alibaba.wt.order.exchange.partner.checktbuser

积分兑换之前校验淘宝账号是否存在
*/
type AlibabaWtOrderExchangePartnerChecktbuserRequest struct {
    model.Params

    // model入参
    outExchangeModel   *OutExchangeModel 

}

func NewAlibabaWtOrderExchangePartnerChecktbuserRequest() *AlibabaWtOrderExchangePartnerChecktbuserRequest{
    return &AlibabaWtOrderExchangePartnerChecktbuserRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetApiMethodName() string {
    return "alibaba.wt.order.exchange.partner.checktbuser"
}

func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWtOrderExchangePartnerChecktbuserRequest) SetOutExchangeModel(outExchangeModel *OutExchangeModel) error {
    r.outExchangeModel = outExchangeModel
    r.Set("out_exchange_model", outExchangeModel)
    return nil
}

func (r AlibabaWtOrderExchangePartnerChecktbuserRequest) GetOutExchangeModel() *OutExchangeModel {
    return r.outExchangeModel
}

