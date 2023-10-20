package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaCfoIncomingInvoiceLedgerFullysync 票易通全量底账数据同步
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
func AlibabaCfoIncomingInvoiceLedgerFullysync(clt *core.SDKClient, req *fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIRequest, resp *fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
