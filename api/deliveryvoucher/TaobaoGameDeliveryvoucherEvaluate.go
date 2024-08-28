package deliveryvoucher

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherEvaluate 卡券评价回传
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
func TaobaoGameDeliveryvoucherEvaluate(ctx context.Context, clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherEvaluateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
