package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceFlowRefund 退订工单(入驻、加盘、续约)
// alibaba.einvoice.flow.refund
//
// 电子发票工单系统，工单退订能力开放
func AlibabaEinvoiceFlowRefund(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceFlowRefundAPIRequest, resp *einvoice.AlibabaEinvoiceFlowRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
