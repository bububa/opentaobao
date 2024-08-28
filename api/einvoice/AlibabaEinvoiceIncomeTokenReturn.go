package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeTokenReturn 服务商回传税号token
// alibaba.einvoice.income.token.return
//
// 服务商回传税号token，用来勾选抵扣认证
func AlibabaEinvoiceIncomeTokenReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeTokenReturnAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeTokenReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
