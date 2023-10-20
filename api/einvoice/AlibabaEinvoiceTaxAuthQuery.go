package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxAuthQuery 发票中台授权信息获取
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
func AlibabaEinvoiceTaxAuthQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxAuthQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxAuthQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
