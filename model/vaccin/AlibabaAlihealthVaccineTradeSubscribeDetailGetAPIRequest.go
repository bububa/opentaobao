package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest 私立疫苗交易-预约详情获取 API请求
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
type AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest struct {
	model.Params
	// 请求参数类
	_tradeSubscribeDetailQueryTopRequest *TradeSubscribeDetailQueryTopRequest
}

// NewAlibabaalihealthvaccinetradesubscribedetailgetRequest 初始化AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest对象
func NewAlibabaalihealthvaccinetradesubscribedetailgetRequest() *AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest {
	return &AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.subscribe.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeSubscribeDetailQueryTopRequest is TradeSubscribeDetailQueryTopRequest Setter
// 请求参数类
func (r *AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest) SetTradeSubscribeDetailQueryTopRequest(_tradeSubscribeDetailQueryTopRequest *TradeSubscribeDetailQueryTopRequest) error {
	r._tradeSubscribeDetailQueryTopRequest = _tradeSubscribeDetailQueryTopRequest
	r.Set("trade_subscribe_detail_query_top_request", _tradeSubscribeDetailQueryTopRequest)
	return nil
}

// GetTradeSubscribeDetailQueryTopRequest TradeSubscribeDetailQueryTopRequest Getter
func (r AlibabaalihealthvaccinetradesubscribedetailgetAPIRequest) GetTradeSubscribeDetailQueryTopRequest() *TradeSubscribeDetailQueryTopRequest {
	return r._tradeSubscribeDetailQueryTopRequest
}
