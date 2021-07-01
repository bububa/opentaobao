package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

/* AlibabaTaxInvoiceSyncLedger
同步底账数据
alibaba.tax.invoice.sync.ledger

接收第三方服务（如：票易通）同步过来的底账发票数据 */
func AlibabaTaxInvoiceSyncLedger(clt *core.SDKClient, req *fundplatform.AlibabaTaxInvoiceSyncLedgerAPIRequest, session string) (*fundplatform.AlibabaTaxInvoiceSyncLedgerAPIResponse, error) {
	var resp fundplatform.AlibabaTaxInvoiceSyncLedgerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
