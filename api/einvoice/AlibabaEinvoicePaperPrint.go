package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePaperPrint 纸票打印接口
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
func AlibabaEinvoicePaperPrint(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperPrintAPIRequest, resp *einvoice.AlibabaEinvoicePaperPrintAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
