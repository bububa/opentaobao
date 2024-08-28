package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundDetailGet 退款详情页渲染
// taobao.refund.detail.get
//
// 退款详情页渲染
func TaobaoRefundDetailGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundDetailGetAPIRequest, resp *tbrefund.TaobaoRefundDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
