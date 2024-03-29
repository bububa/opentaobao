package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest 私立疫苗交易-预约详情更新或保存 API请求
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
type AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest struct {
	model.Params
	// 入参
	_tradeSubscribeDetailExecuteTopRequest *TradeSubscribeDetailExecuteTopRequest
}

// NewAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest 初始化AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest对象
func NewAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest() *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest {
	return &AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) Reset() {
	r._tradeSubscribeDetailExecuteTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.vaccine.trade.subscribe.detail.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeSubscribeDetailExecuteTopRequest is TradeSubscribeDetailExecuteTopRequest Setter
// 入参
func (r *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) SetTradeSubscribeDetailExecuteTopRequest(_tradeSubscribeDetailExecuteTopRequest *TradeSubscribeDetailExecuteTopRequest) error {
	r._tradeSubscribeDetailExecuteTopRequest = _tradeSubscribeDetailExecuteTopRequest
	r.Set("trade_subscribe_detail_execute_top_request", _tradeSubscribeDetailExecuteTopRequest)
	return nil
}

// GetTradeSubscribeDetailExecuteTopRequest TradeSubscribeDetailExecuteTopRequest Getter
func (r AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) GetTradeSubscribeDetailExecuteTopRequest() *TradeSubscribeDetailExecuteTopRequest {
	return r._tradeSubscribeDetailExecuteTopRequest
}

var poolAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest()
	},
}

// GetAlibabaAlihealthVaccineTradeSubscribeDetailSaveRequest 从 sync.Pool 获取 AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest
func GetAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest() *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest {
	return poolAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest.Get().(*AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest)
}

// ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest 将 AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest(v *AlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthVaccineTradeSubscribeDetailSaveAPIRequest.Put(v)
}
