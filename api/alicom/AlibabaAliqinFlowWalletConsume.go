package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletConsume 流量扣减
// alibaba.aliqin.flow.wallet.consume
//
// 流量钱包流量扣减接口
func AlibabaAliqinFlowWalletConsume(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletConsumeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
