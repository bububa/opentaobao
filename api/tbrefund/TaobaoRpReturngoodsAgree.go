package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsAgree 卖家同意退货
// taobao.rp.returngoods.agree
//
// 卖家同意退货，支持淘宝和天猫的订单。
func TaobaoRpReturngoodsAgree(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsAgreeAPIRequest, resp *tbrefund.TaobaoRpReturngoodsAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
