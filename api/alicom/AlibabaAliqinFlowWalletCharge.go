package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletCharge 流量直充
// alibaba.aliqin.flow.wallet.charge
//
// 流量直充
func AlibabaAliqinFlowWalletCharge(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletChargeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletChargeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
