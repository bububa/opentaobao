package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商积分兑换接口 API请求
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换
*/
type AlibabaAlicomOrderExchangeCreateRequest struct {
    model.Params
    // 入参
    _exchangeModel   *ExchangeModel
}

// 初始化AlibabaAlicomOrderExchangeCreateRequest对象
func NewAlibabaAlicomOrderExchangeCreateRequest() *AlibabaAlicomOrderExchangeCreateRequest{
    return &AlibabaAlicomOrderExchangeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderExchangeCreateRequest) GetApiMethodName() string {
    return "alibaba.alicom.order.exchange.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderExchangeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExchangeModel Setter
// 入参
func (r *AlibabaAlicomOrderExchangeCreateRequest) SetExchangeModel(_exchangeModel *ExchangeModel) error {
    r._exchangeModel = _exchangeModel
    r.Set("exchange_model", _exchangeModel)
    return nil
}

// ExchangeModel Getter
func (r AlibabaAlicomOrderExchangeCreateRequest) GetExchangeModel() *ExchangeModel {
    return r._exchangeModel
}
