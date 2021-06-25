package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口tae APIRequest
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
type AlibabaAlicomExchangeCreatebymixnickRequest struct {
    model.Params

    // 入参
    exchangeModel   *ExchangeModel 

}

func NewAlibabaAlicomExchangeCreatebymixnickRequest() *AlibabaAlicomExchangeCreatebymixnickRequest{
    return &AlibabaAlicomExchangeCreatebymixnickRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetApiMethodName() string {
    return "alibaba.alicom.exchange.createbymixnick"
}

func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomExchangeCreatebymixnickRequest) SetExchangeModel(exchangeModel *ExchangeModel) error {
    r.exchangeModel = exchangeModel
    r.Set("exchange_model", exchangeModel)
    return nil
}

func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetExchangeModel() *ExchangeModel {
    return r.exchangeModel
}

