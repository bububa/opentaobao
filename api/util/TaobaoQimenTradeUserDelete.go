package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoQimenTradeUserDelete 删除奇门订单链路用户
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
func TaobaoQimenTradeUserDelete(clt *core.SDKClient, req *util.TaobaoQimenTradeUserDeleteAPIRequest, resp *util.TaobaoQimenTradeUserDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
