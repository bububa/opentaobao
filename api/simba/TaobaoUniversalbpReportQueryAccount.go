package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportqueryaccount 账户报表查询
// taobao.universalbp.report.query.account
//
// 账户报表查询
func Taobaouniversalbpreportqueryaccount(clt *core.SDKClient, req *simba.TaobaouniversalbpreportqueryaccountAPIRequest, session string) (*simba.TaobaouniversalbpreportqueryaccountAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportqueryaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
