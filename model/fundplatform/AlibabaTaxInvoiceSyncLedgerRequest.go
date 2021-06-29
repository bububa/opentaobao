package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步底账数据 APIRequest
alibaba.tax.invoice.sync.ledger

接收第三方服务（如：票易通）同步过来的底账发票数据
*/
type AlibabaTaxInvoiceSyncLedgerRequest struct {
    model.Params

    // 参数
    paramSyncLedgerInvoiceRequest   *SyncLedgerInvoiceRequest 

}

func NewAlibabaTaxInvoiceSyncLedgerRequest() *AlibabaTaxInvoiceSyncLedgerRequest{
    return &AlibabaTaxInvoiceSyncLedgerRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTaxInvoiceSyncLedgerRequest) GetApiMethodName() string {
    return "alibaba.tax.invoice.sync.ledger"
}

func (r AlibabaTaxInvoiceSyncLedgerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTaxInvoiceSyncLedgerRequest) SetParamSyncLedgerInvoiceRequest(paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest) error {
    r.paramSyncLedgerInvoiceRequest = paramSyncLedgerInvoiceRequest
    r.Set("param_sync_ledger_invoice_request", paramSyncLedgerInvoiceRequest)
    return nil
}

func (r AlibabaTaxInvoiceSyncLedgerRequest) GetParamSyncLedgerInvoiceRequest() *SyncLedgerInvoiceRequest {
    return r.paramSyncLedgerInvoiceRequest
}

