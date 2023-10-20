package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceCreateResultsIncrementGet ERP增量开票结果获取
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
func AlibabaEinvoiceCreateResultsIncrementGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreateResultsIncrementGetAPIRequest, resp *einvoice.AlibabaEinvoiceCreateResultsIncrementGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
