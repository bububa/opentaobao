package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceBillEinvoiceList 扫码开票列表
// alibaba.einvoice.bill.einvoice.list
//
// 扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据
func AlibabaEinvoiceBillEinvoiceList(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceBillEinvoiceListAPIRequest, resp *einvoice.AlibabaEinvoiceBillEinvoiceListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
