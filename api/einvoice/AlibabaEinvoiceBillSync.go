package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceBillSync 结算单同步
// alibaba.einvoice.bill.sync
//
// 电子发票业务，服务商同步结算单，包括结算单的增删改功能。最终用于开发票
func AlibabaEinvoiceBillSync(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceBillSyncAPIRequest, resp *einvoice.AlibabaEinvoiceBillSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
