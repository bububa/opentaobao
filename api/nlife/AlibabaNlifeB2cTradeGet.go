package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradeGet 零售+平台查询订单
// alibaba.nlife.b2c.trade.get
//
// 查询零售+平台创建出来的订单详情
func AlibabaNlifeB2cTradeGet(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeGetAPIRequest, resp *nlife.AlibabaNlifeB2cTradeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
