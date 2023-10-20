package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAccountGetBalance 获取账户余额，现金余额
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
func TaobaoUniversalbpAccountGetBalance(clt *core.SDKClient, req *simba.TaobaoUniversalbpAccountGetBalanceAPIRequest, resp *simba.TaobaoUniversalbpAccountGetBalanceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
