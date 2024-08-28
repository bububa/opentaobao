package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoAlitripTravelTradeMemoUpdate 修改一笔交易备注
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
func TaobaoAlitripTravelTradeMemoUpdate(ctx context.Context, clt *core.SDKClient, req *traveltrade.TaobaoAlitripTravelTradeMemoUpdateAPIRequest, resp *traveltrade.TaobaoAlitripTravelTradeMemoUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
