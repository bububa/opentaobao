package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceBillForwordCreate 扫码开票结算单同步前开发票
// alibaba.einvoice.bill.forword.create
//
// 扫码开票结算单同步前开发票，会将数据同步到结算单中
func AlibabaEinvoiceBillForwordCreate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceBillForwordCreateAPIRequest, resp *einvoice.AlibabaEinvoiceBillForwordCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
