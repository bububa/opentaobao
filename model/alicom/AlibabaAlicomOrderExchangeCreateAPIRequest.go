package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderExchangeCreateAPIRequest
代理商积分兑换接口 API请求
alibaba.alicom.order.exchange.create

代理商调用该接口来进行积分兑换 */
type AlibabaAlicomOrderExchangeCreateAPIRequest struct {
	model.Params
	// 入参
	_exchangeModel *ExchangeModel
}

// NewAlibabaAlicomOrderExchangeCreateRequest 初始化AlibabaAlicomOrderExchangeCreateAPIRequest对象
func NewAlibabaAlicomOrderExchangeCreateRequest() *AlibabaAlicomOrderExchangeCreateAPIRequest {
	return &AlibabaAlicomOrderExchangeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomOrderExchangeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.order.exchange.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomOrderExchangeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExchangeModel Setter
// 入参
func (r *AlibabaAlicomOrderExchangeCreateAPIRequest) SetExchangeModel(_exchangeModel *ExchangeModel) error {
	r._exchangeModel = _exchangeModel
	r.Set("exchange_model", _exchangeModel)
	return nil
}

// Get ExchangeModel Getter
func (r AlibabaAlicomOrderExchangeCreateAPIRequest) GetExchangeModel() *ExchangeModel {
	return r._exchangeModel
}
