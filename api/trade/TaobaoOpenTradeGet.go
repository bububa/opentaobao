package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoOpenTradeGet 获取单笔交易的部分信息(商家应用使用)
// taobao.open.trade.get
//
// 获取单笔交易的部分信息&lt;/br&gt;
// 1.入参fields中传入buyer_nick ，才能返回buyer_open_id
func TaobaoOpenTradeGet(clt *core.SDKClient, req *trade.TaobaoOpenTradeGetAPIRequest, resp *trade.TaobaoOpenTradeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
