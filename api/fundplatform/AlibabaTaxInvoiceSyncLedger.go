package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaTaxInvoiceSyncLedger 同步底账数据
// alibaba.tax.invoice.sync.ledger
//
// 接收第三方服务（如：票易通）同步过来的底账发票数据
func AlibabaTaxInvoiceSyncLedger(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaTaxInvoiceSyncLedgerAPIRequest, resp *fundplatform.AlibabaTaxInvoiceSyncLedgerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
