package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelTradeClose 飞猪度假-订单关闭接口（快速退款）
// alitrip.travel.trade.close
//
// 卖家关单（快速退款接口），不支持二次预约商品的订单
func AlitripTravelTradeClose(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelTradeCloseAPIRequest, resp *traveltrade.AlitripTravelTradeCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
