package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePaperPrint 纸票打印接口
// alibaba.einvoice.paper.print
//
// 打印一张已开具成功的纸票
func AlibabaEinvoicePaperPrint(clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperPrintAPIRequest, resp *einvoice.AlibabaEinvoicePaperPrintAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
