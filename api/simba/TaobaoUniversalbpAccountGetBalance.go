package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpaccountgetbalance 获取账户余额，现金余额
// taobao.universalbp.account.get.balance
//
// 获取账户实时现金余额
func Taobaouniversalbpaccountgetbalance(clt *core.SDKClient, req *simba.TaobaouniversalbpaccountgetbalanceAPIRequest, session string) (*simba.TaobaouniversalbpaccountgetbalanceAPIResponse, error) {
	var resp simba.TaobaouniversalbpaccountgetbalanceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
