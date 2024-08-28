package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcTradesBytagGet 标签查询订单
// taobao.oc.trades.bytag.get
//
// 根据标签查询订单编号
func TaobaoOcTradesBytagGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoOcTradesBytagGetAPIRequest, resp *jst.TaobaoOcTradesBytagGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
