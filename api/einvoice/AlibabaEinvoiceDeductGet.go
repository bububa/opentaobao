package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceDeductGet 发票扣减的接口
// alibaba.einvoice.deduct.get
//
// 获取历史发票扣减量、每日发票扣减量的接口
func AlibabaEinvoiceDeductGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceDeductGetAPIRequest, resp *einvoice.AlibabaEinvoiceDeductGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
