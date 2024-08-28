package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQueryauthstate 查询连包签约状态
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
func YoukuOttPayOrderQueryauthstate(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryauthstateAPIRequest, resp *ottpay.YoukuOttPayOrderQueryauthstateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
