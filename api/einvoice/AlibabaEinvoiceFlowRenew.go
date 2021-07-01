package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

/* AlibabaEinvoiceFlowRenew
工单(入驻、加盘、续约)续约
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放 */
func AlibabaEinvoiceFlowRenew(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceFlowRenewAPIRequest, session string) (*einvoice.AlibabaEinvoiceFlowRenewAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceFlowRenewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
