package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest 私立疫苗交易-预约详情更新或保存 API请求
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
type AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest struct {
	model.Params
	// 入参
	_tradeSubscribeDetailExecuteTopRequest *TradeSubscribeDetailExecuteTopRequest
}

// NewAlibabaalihealthvaccinetradesubscribedetailsaveRequest 初始化AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest对象
func NewAlibabaalihealthvaccinetradesubscribedetailsaveRequest() *AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest {
	return &AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.subscribe.detail.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeSubscribeDetailExecuteTopRequest is TradeSubscribeDetailExecuteTopRequest Setter
// 入参
func (r *AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest) SetTradeSubscribeDetailExecuteTopRequest(_tradeSubscribeDetailExecuteTopRequest *TradeSubscribeDetailExecuteTopRequest) error {
	r._tradeSubscribeDetailExecuteTopRequest = _tradeSubscribeDetailExecuteTopRequest
	r.Set("trade_subscribe_detail_execute_top_request", _tradeSubscribeDetailExecuteTopRequest)
	return nil
}

// GetTradeSubscribeDetailExecuteTopRequest TradeSubscribeDetailExecuteTopRequest Getter
func (r AlibabaalihealthvaccinetradesubscribedetailsaveAPIRequest) GetTradeSubscribeDetailExecuteTopRequest() *TradeSubscribeDetailExecuteTopRequest {
	return r._tradeSubscribeDetailExecuteTopRequest
}
