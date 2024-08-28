package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQueryorder 查询订单
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
func YoukuOttPayOrderQueryorder(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryorderAPIRequest, resp *ottpay.YoukuOttPayOrderQueryorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
