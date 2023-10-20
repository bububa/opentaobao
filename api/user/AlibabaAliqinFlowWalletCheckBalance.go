package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaAliqinFlowWalletCheckBalance 商家预存余额检查
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
func AlibabaAliqinFlowWalletCheckBalance(clt *core.SDKClient, req *user.AlibabaAliqinFlowWalletCheckBalanceAPIRequest, resp *user.AlibabaAliqinFlowWalletCheckBalanceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
