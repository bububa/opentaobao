package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoQimenTradeUserAdd 添加奇门订单链路用户
// taobao.qimen.trade.user.add
//
// 添加奇门订单链路用户
func TaobaoQimenTradeUserAdd(ctx context.Context, clt *core.SDKClient, req *util.TaobaoQimenTradeUserAddAPIRequest, resp *util.TaobaoQimenTradeUserAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
