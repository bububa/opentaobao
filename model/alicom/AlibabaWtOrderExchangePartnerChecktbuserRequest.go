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
type AlibabaWtOrderExchangePartnerChecktbuserAPIRequest struct {
    model.Params
    // model入参
    _outExchangeModel   *OutExchangeModel
}

// 初始化AlibabaWtOrderExchangePartnerChecktbuserAPIRequest对象
func NewAlibabaWtOrderExchangePartnerChecktbuserRequest() *AlibabaWtOrderExchangePartnerChecktbuserAPIRequest{
    return &AlibabaWtOrderExchangePartnerChecktbuserAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWtOrderExchangePartnerChecktbuserAPIRequest) GetApiMethodName() string {
    return "alibaba.wt.order.exchange.partner.checktbuser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWtOrderExchangePartnerChecktbuserAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutExchangeModel Setter
// model入参
func (r *AlibabaWtOrderExchangePartnerChecktbuserAPIRequest) SetOutExchangeModel(_outExchangeModel *OutExchangeModel) error {
    r._outExchangeModel = _outExchangeModel
    r.Set("out_exchange_model", _outExchangeModel)
    return nil
}

// OutExchangeModel Getter
func (r AlibabaWtOrderExchangePartnerChecktbuserAPIRequest) GetOutExchangeModel() *OutExchangeModel {
    return r._outExchangeModel
}
