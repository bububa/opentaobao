package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest 私立疫苗交易-预约详情获取 API请求
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
type AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest struct {
	model.Params
	// 请求参数类
	_tradeSubscribeDetailQueryTopRequest *TradeSubscribeDetailQueryTopRequest
}

// NewAlibabaAlihealthVaccineTradeSubscribeDetailGetRequest 初始化AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest对象
func NewAlibabaAlihealthVaccineTradeSubscribeDetailGetRequest() *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest {
	return &AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) Reset() {
	r._tradeSubscribeDetailQueryTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.subscribe.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeSubscribeDetailQueryTopRequest is TradeSubscribeDetailQueryTopRequest Setter
// 请求参数类
func (r *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) SetTradeSubscribeDetailQueryTopRequest(_tradeSubscribeDetailQueryTopRequest *TradeSubscribeDetailQueryTopRequest) error {
	r._tradeSubscribeDetailQueryTopRequest = _tradeSubscribeDetailQueryTopRequest
	r.Set("trade_subscribe_detail_query_top_request", _tradeSubscribeDetailQueryTopRequest)
	return nil
}

// GetTradeSubscribeDetailQueryTopRequest TradeSubscribeDetailQueryTopRequest Getter
func (r AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) GetTradeSubscribeDetailQueryTopRequest() *TradeSubscribeDetailQueryTopRequest {
	return r._tradeSubscribeDetailQueryTopRequest
}

var poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthVaccineTradeSubscribeDetailGetRequest()
	},
}

// GetAlibabaAlihealthVaccineTradeSubscribeDetailGetRequest 从 sync.Pool 获取 AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest
func GetAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest() *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest {
	return poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest.Get().(*AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest)
}

// ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest 将 AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest(v *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIRequest.Put(v)
}
