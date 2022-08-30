package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest 通过订单ID与卖家ID获取订单渠道 API请求
// alibaba.alihealth.vaccine.trade.order.channel.get
//
// 通过订单ID与卖家ID获取订单渠道
type AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest struct {
	model.Params
	// 请求参数
	_tradeVaccineOrderQueryTopRequest *TradeVaccineOrderQueryTopRequest
}

// NewAlibabaAlihealthVaccineTradeOrderChannelGetRequest 初始化AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest对象
func NewAlibabaAlihealthVaccineTradeOrderChannelGetRequest() *AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest {
	return &AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.order.channel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeVaccineOrderQueryTopRequest is TradeVaccineOrderQueryTopRequest Setter
// 请求参数
func (r *AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) SetTradeVaccineOrderQueryTopRequest(_tradeVaccineOrderQueryTopRequest *TradeVaccineOrderQueryTopRequest) error {
	r._tradeVaccineOrderQueryTopRequest = _tradeVaccineOrderQueryTopRequest
	r.Set("trade_vaccine_order_query_top_request", _tradeVaccineOrderQueryTopRequest)
	return nil
}

// GetTradeVaccineOrderQueryTopRequest TradeVaccineOrderQueryTopRequest Getter
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetTradeVaccineOrderQueryTopRequest() *TradeVaccineOrderQueryTopRequest {
	return r._tradeVaccineOrderQueryTopRequest
}
