package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAccountGetBalance 获取账户余额，现金余额
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
func TaobaoUniversalbpAccountGetBalance(clt *core.SDKClient, req *simba.TaobaoUniversalbpAccountGetBalanceAPIRequest, session string) (*simba.TaobaoUniversalbpAccountGetBalanceAPIResponse, error) {
	var resp simba.TaobaoUniversalbpAccountGetBalanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
