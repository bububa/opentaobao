package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceOrderRefundUpdate 回传订单退款审核结果
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
func AlibabaEinvoiceOrderRefundUpdate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceOrderRefundUpdateAPIRequest, resp *einvoice.AlibabaEinvoiceOrderRefundUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
