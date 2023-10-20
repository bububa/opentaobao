package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletChargeRule 流量钱包直充（根据号码归属地省份路由）
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
func AlibabaAliqinFlowWalletChargeRule(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletChargeRuleAPIRequest, session string) (*alicom.AlibabaAliqinFlowWalletChargeRuleAPIResponse, error) {
	var resp alicom.AlibabaAliqinFlowWalletChargeRuleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
