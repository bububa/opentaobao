package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersList 全渠道经销商订单列表
// taobao.omni.dealer.oders.list
//
// 全渠道经销商订单列表查询
func TaobaoOmniDealerOdersList(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersListAPIRequest, resp *omniorder.TaobaoOmniDealerOdersListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
