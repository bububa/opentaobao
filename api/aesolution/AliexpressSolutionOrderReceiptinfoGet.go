package aesolution

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aesolution"
)

// AliexpressSolutionOrderReceiptinfoGet Get Order Receipt Info
// aliexpress.solution.order.receiptinfo.get
//
// Get Order Receipt Info, Support multi stores requirements for Turkey sellers.
func AliexpressSolutionOrderReceiptinfoGet(ctx context.Context, clt *core.SDKClient, req *aesolution.AliexpressSolutionOrderReceiptinfoGetAPIRequest, resp *aesolution.AliexpressSolutionOrderReceiptinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
