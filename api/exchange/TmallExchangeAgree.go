package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeAgree 卖家同意换货申请
// tmall.exchange.agree
//
// 卖家同意换货申请
func TmallExchangeAgree(clt *core.SDKClient, req *exchange.TmallExchangeAgreeAPIRequest, resp *exchange.TmallExchangeAgreeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
