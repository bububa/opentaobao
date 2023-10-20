package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaxInvoiceSyncLedgerAPIRequest 同步底账数据 API请求
// alibaba.tax.invoice.sync.ledger
//
// 接收第三方服务（如：票易通）同步过来的底账发票数据
type AlibabaTaxInvoiceSyncLedgerAPIRequest struct {
	model.Params
	// 参数
	_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest
}

// NewAlibabaTaxInvoiceSyncLedgerRequest 初始化AlibabaTaxInvoiceSyncLedgerAPIRequest对象
func NewAlibabaTaxInvoiceSyncLedgerRequest() *AlibabaTaxInvoiceSyncLedgerAPIRequest {
	return &AlibabaTaxInvoiceSyncLedgerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaxInvoiceSyncLedgerAPIRequest) Reset() {
	r._paramSyncLedgerInvoiceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaxInvoiceSyncLedgerAPIRequest) GetApiMethodName() string {
	return "alibaba.tax.invoice.sync.ledger"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaxInvoiceSyncLedgerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaxInvoiceSyncLedgerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamSyncLedgerInvoiceRequest is ParamSyncLedgerInvoiceRequest Setter
// 参数
func (r *AlibabaTaxInvoiceSyncLedgerAPIRequest) SetParamSyncLedgerInvoiceRequest(_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest) error {
	r._paramSyncLedgerInvoiceRequest = _paramSyncLedgerInvoiceRequest
	r.Set("param_sync_ledger_invoice_request", _paramSyncLedgerInvoiceRequest)
	return nil
}

// GetParamSyncLedgerInvoiceRequest ParamSyncLedgerInvoiceRequest Getter
func (r AlibabaTaxInvoiceSyncLedgerAPIRequest) GetParamSyncLedgerInvoiceRequest() *SyncLedgerInvoiceRequest {
	return r._paramSyncLedgerInvoiceRequest
}

var poolAlibabaTaxInvoiceSyncLedgerAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaxInvoiceSyncLedgerRequest()
	},
}

// GetAlibabaTaxInvoiceSyncLedgerRequest 从 sync.Pool 获取 AlibabaTaxInvoiceSyncLedgerAPIRequest
func GetAlibabaTaxInvoiceSyncLedgerAPIRequest() *AlibabaTaxInvoiceSyncLedgerAPIRequest {
	return poolAlibabaTaxInvoiceSyncLedgerAPIRequest.Get().(*AlibabaTaxInvoiceSyncLedgerAPIRequest)
}

// ReleaseAlibabaTaxInvoiceSyncLedgerAPIRequest 将 AlibabaTaxInvoiceSyncLedgerAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaxInvoiceSyncLedgerAPIRequest(v *AlibabaTaxInvoiceSyncLedgerAPIRequest) {
	v.Reset()
	poolAlibabaTaxInvoiceSyncLedgerAPIRequest.Put(v)
}
