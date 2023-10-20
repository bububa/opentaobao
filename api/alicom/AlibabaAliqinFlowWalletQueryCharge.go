package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletQueryCharge 查询流量充值状态
// alibaba.aliqin.flow.wallet.query.charge
//
// 查询流量充值状态
func AlibabaAliqinFlowWalletQueryCharge(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletQueryChargeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletQueryChargeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
