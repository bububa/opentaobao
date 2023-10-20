package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantResultGet 商家自研ERP开票结果获取
// alibaba.einvoice.merchant.result.get
//
// 商家自研ERP开票结果获取
func AlibabaEinvoiceMerchantResultGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantResultGetAPIRequest, resp *einvoice.AlibabaEinvoiceMerchantResultGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
