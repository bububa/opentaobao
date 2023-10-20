package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaCfoIncomingInvoiceLedgerFullysync 票易通全量底账数据同步
// alibaba.cfo.incoming.invoice.ledger.fullysync
//
// 票易通全量底账数据同步
func AlibabaCfoIncomingInvoiceLedgerFullysync(clt *core.SDKClient, req *fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIRequest, session string) (*fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse, error) {
	var resp fundplatform.AlibabaCfoIncomingInvoiceLedgerFullysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
