package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptCustGet 获取账户实时报表数据
// taobao.simba.rtrpt.cust.get
//
// 获取账户实时报表数据
func TaobaoSimbaRtrptCustGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptCustGetAPIRequest, resp *simba.TaobaoSimbaRtrptCustGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
