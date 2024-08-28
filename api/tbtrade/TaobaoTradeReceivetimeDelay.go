package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeReceivetimeDelay 延长交易收货时间
// taobao.trade.receivetime.delay
//
// 延长交易收货时间
func TaobaoTradeReceivetimeDelay(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTradeReceivetimeDelayAPIRequest, resp *tbtrade.TaobaoTradeReceivetimeDelayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
