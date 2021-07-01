package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTaxInvoiceSyncLedgerAPIRequest
同步底账数据 API请求
alibaba.tax.invoice.sync.ledger

接收第三方服务（如：票易通）同步过来的底账发票数据 */
type AlibabaTaxInvoiceSyncLedgerAPIRequest struct {
	model.Params
	// 参数
	_paramSyncLedgerInvoiceRequest *SyncLedgerInvoiceRequest
}

// New
