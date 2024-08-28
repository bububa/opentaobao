package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletChargeRule 流量钱包直充（根据号码归属地省份路由）
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
func AlibabaAliqinFlowWalletChargeRule(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletChargeRuleAPIRequest, resp *alicom.AlibabaAliqinFlowWalletChargeRuleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
