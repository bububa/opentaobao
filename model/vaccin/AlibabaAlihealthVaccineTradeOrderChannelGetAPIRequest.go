package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccinetradeorderchannelgetAPIRequest 通过订单ID与卖家ID获取订单渠道 API请求
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
type AlibabaalihealthvaccinetradeorderchannelgetAPIRequest struct {
	model.Params
	// 请求参数
	_tradeVaccineOrderQueryTopRequest *TradeVaccineOrderQueryTopRequest
}

// NewAlibabaalihealthvaccinetradeorderchannelgetRequest 初始化AlibabaalihealthvaccinetradeorderchannelgetAPIRequest对象
func NewAlibabaalihealthvaccinetradeorderchannelgetRequest() *AlibabaalihealthvaccinetradeorderchannelgetAPIRequest {
	return &AlibabaalihealthvaccinetradeorderchannelgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthvaccinetradeorderchannelgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.order.channel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthvaccinetradeorderchannelgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthvaccinetradeorderchannelgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeVaccineOrderQueryTopRequest is TradeVaccineOrderQueryTopRequest Setter
// 请求参数
func (r *AlibabaalihealthvaccinetradeorderchannelgetAPIRequest) SetTradeVaccineOrderQueryTopRequest(_tradeVaccineOrderQueryTopRequest *TradeVaccineOrderQueryTopRequest) error {
	r._tradeVaccineOrderQueryTopRequest = _tradeVaccineOrderQueryTopRequest
	r.Set("trade_vaccine_order_query_top_request", _tradeVaccineOrderQueryTopRequest)
	return nil
}

// GetTradeVaccineOrderQueryTopRequest TradeVaccineOrderQueryTopRequest Getter
func (r AlibabaalihealthvaccinetradeorderchannelgetAPIRequest) GetTradeVaccineOrderQueryTopRequest() *TradeVaccineOrderQueryTopRequest {
	return r._tradeVaccineOrderQueryTopRequest
}
