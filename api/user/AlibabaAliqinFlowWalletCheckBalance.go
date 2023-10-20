package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibabaaliqinflowwalletcheckbalance 商家预存余额检查
// alibaba.aliqin.flow.wallet.check.balance
//
// 检查商家CRM预存余额是否足够进行活动
func Alibabaaliqinflowwalletcheckbalance(clt *core.SDKClient, req *user.AlibabaaliqinflowwalletcheckbalanceAPIRequest, session string) (*user.AlibabaaliqinflowwalletcheckbalanceAPIResponse, error) {
	var resp user.AlibabaaliqinflowwalletcheckbalanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
