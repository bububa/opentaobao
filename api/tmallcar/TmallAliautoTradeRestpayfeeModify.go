package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeRestpayfeeModify 分阶段订单尾款改价
// tmall.aliauto.trade.restpayfee.modify
//
// 汽车商家通过此api修改整车分阶段订单的尾款金额
func TmallAliautoTradeRestpayfeeModify(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoTradeRestpayfeeModifyAPIRequest, resp *tmallcar.TmallAliautoTradeRestpayfeeModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
