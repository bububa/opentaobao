package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQuerycporder 查询支付订单对应cp订单号
// youku.ott.pay.order.querycporder
//
// 根据支付订单查询对应cp订单号
func YoukuOttPayOrderQuerycporder(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQuerycporderAPIRequest, resp *ottpay.YoukuOttPayOrderQuerycporderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
