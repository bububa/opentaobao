package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomExchangeCreatebymixnickAPIRequest
代理商积分兑换接口tae API请求
alibaba.alicom.exchange.createbymixnick

代理商调用该接口来进行积分兑换，tae */
type AlibabaAlicomExchangeCreatebymixnickAPIRequest struct {
	model.Params
	// 入参
	_exchangeModel *ExchangeModel
}

// NewAlibabaAlicomExchangeCreatebymixnickRequest 初始化AlibabaAlicomExchangeCreatebymixnickAPIRequest对象
func NewAlibabaAlicomExchangeCreatebymixnickRequest() *AlibabaAlicomExchangeCreatebymixnickAPIRequest {
	return &AlibabaAlicomExchangeCreatebymixnickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlicomExchangeCreatebymixnickAPIRequest) GetApiMethodName() string {
	return "alibaba.alicom.exchange.createbymixnick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlicomExchangeCreatebymixnickAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExchangeModel Setter
// 入参
func (r *AlibabaAlicomExchangeCreatebymixnickAPIRequest) SetExchangeModel(_exchangeModel *ExchangeModel) error {
	r._exchangeModel = _exchangeModel
	r.Set("exchange_model", _exchangeModel)
	return nil
}

// Get ExchangeModel Getter
func (r AlibabaAlicomExchangeCreatebymixnickAPIRequest) GetExchangeModel() *ExchangeModel {
	return r._exchangeModel
}
