package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersGet 获取单笔全渠道经销商订单的详细信息
// taobao.omni.dealer.oders.get
//
// 全渠道经销商获取单笔订单的详细信息
func TaobaoOmniDealerOdersGet(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersGetAPIRequest, resp *omniorder.TaobaoOmniDealerOdersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
