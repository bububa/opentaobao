package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoLifeReservationTradeConsumeNotice 生服购后预约外部核销
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
func TaobaoLifeReservationTradeConsumeNotice(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoLifeReservationTradeConsumeNoticeAPIRequest, resp *trade.TaobaoLifeReservationTradeConsumeNoticeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
