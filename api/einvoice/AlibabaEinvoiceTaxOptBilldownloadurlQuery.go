package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptBilldownloadurlQuery 税筹业务账单文件下载URL查询
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
func AlibabaEinvoiceTaxOptBilldownloadurlQuery(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptBilldownloadurlQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
