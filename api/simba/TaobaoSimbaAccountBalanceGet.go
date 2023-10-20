package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaaccountbalanceget 获取实时余额，”元”为单位
// taobao.simba.account.balance.get
//
// 获取实时余额，”元”为单位
func Taobaosimbaaccountbalanceget(clt *core.SDKClient, req *simba.TaobaosimbaaccountbalancegetAPIRequest, session string) (*simba.TaobaosimbaaccountbalancegetAPIResponse, error) {
	var resp simba.TaobaosimbaaccountbalancegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
