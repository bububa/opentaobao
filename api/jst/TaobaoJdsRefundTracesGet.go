package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJdsRefundTracesGet 获取单条退款跟踪详情
// taobao.jds.refund.traces.get
//
// 获取聚石塔数据共享的交易全链路的退款信息
func TaobaoJdsRefundTracesGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJdsRefundTracesGetAPIRequest, resp *jst.TaobaoJdsRefundTracesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
