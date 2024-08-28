package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeCarOrderGet 整车订单详情查询
// tmall.aliauto.trade.car.order.get
//
// 整车订单详情查询接口
func TmallAliautoTradeCarOrderGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoTradeCarOrderGetAPIRequest, resp *tmallcar.TmallAliautoTradeCarOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
