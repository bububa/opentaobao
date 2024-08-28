package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRpReturngoodsRefill 卖家回填物流信息
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
func TaobaoRpReturngoodsRefill(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRpReturngoodsRefillAPIRequest, resp *tbrefund.TaobaoRpReturngoodsRefillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
