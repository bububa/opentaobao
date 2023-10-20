package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaAccountBalanceGet 获取实时余额，”元”为单位
// taobao.simba.account.balance.get
//
// 获取实时余额，”元”为单位
func TaobaoSimbaAccountBalanceGet(clt *core.SDKClient, req *simba.TaobaoSimbaAccountBalanceGetAPIRequest, resp *simba.TaobaoSimbaAccountBalanceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
