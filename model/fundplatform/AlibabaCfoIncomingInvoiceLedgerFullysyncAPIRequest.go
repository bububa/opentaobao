package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacfoincominginvoiceledgerfullysyncAPIRequest 票易通全量底账数据同步 API请求
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
type AlibabacfoincominginvoiceledgerfullysyncAPIRequest struct {
	model.Params
	// 底账同步请求
	_paramPytLedgerSyncRequest *PytLedgerSyncRequest
}

// NewAlibabacfoincominginvoiceledgerfullysyncRequest 初始化AlibabacfoincominginvoiceledgerfullysyncAPIRequest对象
func NewAlibabacfoincominginvoiceledgerfullysyncRequest() *AlibabacfoincominginvoiceledgerfullysyncAPIRequest {
	return &AlibabacfoincominginvoiceledgerfullysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacfoincominginvoiceledgerfullysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.cfo.incoming.invoice.ledger.fullysync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacfoincominginvoiceledgerfullysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacfoincominginvoiceledgerfullysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPytLedgerSyncRequest is ParamPytLedgerSyncRequest Setter
// 底账同步请求
func (r *AlibabacfoincominginvoiceledgerfullysyncAPIRequest) SetParamPytLedgerSyncRequest(_paramPytLedgerSyncRequest *PytLedgerSyncRequest) error {
	r._paramPytLedgerSyncRequest = _paramPytLedgerSyncRequest
	r.Set("param_pyt_ledger_sync_request", _paramPytLedgerSyncRequest)
	return nil
}

// GetParamPytLedgerSyncRequest ParamPytLedgerSyncRequest Getter
func (r AlibabacfoincominginvoiceledgerfullysyncAPIRequest) GetParamPytLedgerSyncRequest() *PytLedgerSyncRequest {
	return r._paramPytLedgerSyncRequest
}
