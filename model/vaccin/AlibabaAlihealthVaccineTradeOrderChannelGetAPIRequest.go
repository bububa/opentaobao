package vaccin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) Reset() {
	r._tradeVaccineOrderQueryTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.order.channel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthVaccineTradeOrderChannelGetRequest()
	},
}

// GetAlibabaAlihealthVaccineTradeOrderChannelGetRequest 从 sync.Pool 获取 AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest
func GetAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest() *AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest {
	return poolAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest.Get().(*AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest)
}

// ReleaseAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest 将 AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest(v *AlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthVaccineTradeOrderChannelGetAPIRequest.Put(v)
}
