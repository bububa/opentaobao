package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口tae API请求
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae
*/
type AlibabaAlicomExchangeCreatebymixnickRequest struct {
    model.Params
    // 入参
    _exchangeModel   *ExchangeModel
}

// 初始化AlibabaAlicomExchangeCreatebymixnickRequest对象
func NewAlibabaAlicomExchangeCreatebymixnickRequest() *AlibabaAlicomExchangeCreatebymixnickRequest{
    return &AlibabaAlicomExchangeCreatebymixnickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetApiMethodName() string {
    return "alibaba.alicom.exchange.createbymixnick"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExchangeModel Setter
// 入参
func (r *AlibabaAlicomExchangeCreatebymixnickRequest) SetExchangeModel(_exchangeModel *ExchangeModel) error {
    r._exchangeModel = _exchangeModel
    r.Set("exchange_model", _exchangeModel)
    return nil
}

// ExchangeModel Getter
func (r AlibabaAlicomExchangeCreatebymixnickRequest) GetExchangeModel() *ExchangeModel {
    return r._exchangeModel
}
