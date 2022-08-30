package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest 私立疫苗交易-预约详情更新或保存 API请求
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
type AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest struct {
	model.Params
	// 入参
	_tradeSubscribeDetailSaveTopRequest *TradeSubscribeDetailSaveTopRequest
}

// NewAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest 初始化AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest对象
func NewAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest() *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest {
	return &AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.subscribe.detail.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTradeSubscribeDetailSaveTopRequest is TradeSubscribeDetailSaveTopRequest Setter
// 入参
func (r *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) SetTradeSubscribeDetailSaveTopRequest(_tradeSubscribeDetailSaveTopRequest *TradeSubscribeDetailSaveTopRequest) error {
	r._tradeSubscribeDetailSaveTopRequest = _tradeSubscribeDetailSaveTopRequest
	r.Set("trade_subscribe_detail_save_top_request", _tradeSubscribeDetailSaveTopRequest)
	return nil
}

// GetTradeSubscribeDetailSaveTopRequest TradeSubscribeDetailSaveTopRequest Getter
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetTradeSubscribeDetailSaveTopRequest() *TradeSubscribeDetailSaveTopRequest {
	return r._tradeSubscribeDetailSaveTopRequest
}
