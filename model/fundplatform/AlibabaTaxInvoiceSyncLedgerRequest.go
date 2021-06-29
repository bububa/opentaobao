package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步底账数据 API请求
alibaba.tax.invoice.sync.ledger

接收第三方服务（如：票易通）同步过来的底账发票数据
*/
type AlibabaTaxInvoiceSyncLedgerRequest struct {
    model.Params
    // 参数
    _paramSyncLedgerInvoiceRequest   *SyncLedgerInvoiceRequest
}

// 初始化AlibabaTaxInvoiceSyncLedgerRequest对象
func NewAlibabaTaxInvoiceSyncLedgerRequest() *AlibabaTaxInvoiceSyncLedgerRequest{
    return &AlibabaTaxInvoiceSyncLedgerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTaxInvoiceSyncLedgerRequest) GetApiMethodName() string {
    return "alibaba.tax.invoice.sync.ledger"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTaxInvoiceSyncLedgerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamSyncLedgerInvoiceRequest Setter
// 参数
func (r *AlibabaTaxInvoiceSyncLedgerRequest) SetParamSyncLedgerInvoiceRequest(_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest) error {
    r._paramSyncLedgerInvoiceRequest = _paramSyncLedgerInvoiceRequest
    r.Set("param_sync_ledger_invoice_request", _paramSyncLedgerInvoiceRequest)
    return nil
}

// ParamSyncLedgerInvoiceRequest Getter
func (r AlibabaTaxInvoiceSyncLedgerRequest) GetParamSyncLedgerInvoiceRequest() *SyncLedgerInvoiceRequest {
    return r._paramSyncLedgerInvoiceRequest
}
