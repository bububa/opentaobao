package aesolution

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderReceiptinfoGet Get Order Receipt Info
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
func AliexpressSolutionOrderReceiptinfoGet(clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderReceiptinfoGetAPIRequest, session string) (*aesolution.AliexpressSolutionOrderReceiptinfoGetAPIResponse, error) {
	var resp aesolution.AliexpressSolutionOrderReceiptinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
