package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderCreateorder 创建订单
// youku.ott.pay.order.createorder
//
// ottpay创建订单
func YoukuOttPayOrderCreateorder(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderCreateorderAPIRequest, resp *ottpay.YoukuOttPayOrderCreateorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
