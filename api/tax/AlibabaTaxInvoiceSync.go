package tax

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tax"
)

// AlibabaTaxInvoiceSync 第三方开票回调API
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
func AlibabaTaxInvoiceSync(ctx context.Context, clt *core.SDKClient, req *tax.AlibabaTaxInvoiceSyncAPIRequest, resp *tax.AlibabaTaxInvoiceSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
