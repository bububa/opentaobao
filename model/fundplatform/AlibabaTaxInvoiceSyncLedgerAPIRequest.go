package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabataxinvoicesyncledgerAPIRequest 同步底账数据 API请求
// alibaba.tax.invoice.sync.ledger
//
// 接收第三方服务（如：票易通）同步过来的底账发票数据
type AlibabataxinvoicesyncledgerAPIRequest struct {
	model.Params
	// 参数
	_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest
}

// NewAlibabataxinvoicesyncledgerRequest 初始化AlibabataxinvoicesyncledgerAPIRequest对象
func NewAlibabataxinvoicesyncledgerRequest() *AlibabataxinvoicesyncledgerAPIRequest {
	return &AlibabataxinvoicesyncledgerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabataxinvoicesyncledgerAPIRequest) GetApiMethodName() string {
	return "alibaba.tax.invoice.sync.ledger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabataxinvoicesyncledgerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabataxinvoicesyncledgerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncLedgerInvoiceRequest is ParamSyncLedgerInvoiceRequest Setter
// 参数
func (r *AlibabataxinvoicesyncledgerAPIRequest) SetParamSyncLedgerInvoiceRequest(_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest) error {
	r._paramSyncLedgerInvoiceRequest = _paramSyncLedgerInvoiceRequest
	r.Set("param_sync_ledger_invoice_request", _paramSyncLedgerInvoiceRequest)
	return nil
}

// GetParamSyncLedgerInvoiceRequest ParamSyncLedgerInvoiceRequest Getter
func (r AlibabataxinvoicesyncledgerAPIRequest) GetParamSyncLedgerInvoiceRequest() *SyncLedgerInvoiceRequest {
	return r._paramSyncLedgerInvoiceRequest
}
