package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口 APIRequest
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换
*/
type AlibabaAlicomOrderExchangeCreateRequest struct {
    model.Params

    // 入参
    exchangeModel   *ExchangeModel 

}

func NewAlibabaAlicomOrderExchangeCreateRequest() *AlibabaAlicomOrderExchangeCreateRequest{
    return &AlibabaAlicomOrderExchangeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomOrderExchangeCreateRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.exchange.create"
}

func (r AlibabaAlicomOrderExchangeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomOrderExchangeCreateRequest) SetExchangeModel(exchangeModel *ExchangeModel) error {
    r.exchangeModel = exchangeModel
    r.Set("exchange_model", exchangeModel)
    return nil
}

func (r AlibabaAlicomOrderExchangeCreateRequest) GetExchangeModel() *ExchangeModel {
    return r.exchangeModel
}

