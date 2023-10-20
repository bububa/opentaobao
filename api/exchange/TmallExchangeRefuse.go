package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeRefuse 卖家拒绝换货申请
// tmall.exchange.refuse
//
// 卖家拒绝换货申请
func TmallExchangeRefuse(clt *core.SDKClient, req *exchange.TmallExchangeRefuseAPIRequest, resp *exchange.TmallExchangeRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
