package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRtrptBidwordGet 获取推广词实时报表数据
// taobao.simba.rtrpt.bidword.get
//
// 获取推广词报表数据
func TaobaoSimbaRtrptBidwordGet(clt *core.SDKClient, req *simba.TaobaoSimbaRtrptBidwordGetAPIRequest, resp *simba.TaobaoSimbaRtrptBidwordGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
