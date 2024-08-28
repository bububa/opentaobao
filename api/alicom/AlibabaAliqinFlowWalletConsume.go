package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletConsume 流量扣减
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
func AlibabaAliqinFlowWalletConsume(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletConsumeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
