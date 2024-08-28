package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeVerifyReturn 服务商回传发票查验的结果
// alibaba.einvoice.income.verify.return
//
// 服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传
func AlibabaEinvoiceIncomeVerifyReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeVerifyReturnAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeVerifyReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
