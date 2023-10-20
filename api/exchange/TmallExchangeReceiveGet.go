package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeReceiveGet 卖家查询换货列表
// tmall.exchange.receive.get
//
// 卖家查询换货列表
func TmallExchangeReceiveGet(clt *core.SDKClient, req *exchange.TmallExchangeReceiveGetAPIRequest, resp *exchange.TmallExchangeReceiveGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
