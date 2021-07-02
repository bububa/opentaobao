package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceCreateResultsIncrementGet ERP增量开票结果获取
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
func AlibabaEinvoiceCreateResultsIncrementGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreateResultsIncrementGetAPIRequest, session string) (*einvoice.AlibabaEinvoiceCreateResultsIncrementGetAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceCreateResultsIncrementGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
