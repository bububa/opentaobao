package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniDealerOdersRefundAddress 经销商查询逆向退货地址
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
func TaobaoOmniDealerOdersRefundAddress(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniDealerOdersRefundAddressAPIRequest, resp *omniorder.TaobaoOmniDealerOdersRefundAddressAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
