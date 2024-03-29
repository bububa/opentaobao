package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoLifeReservationTradeConsumeNotice 生服购后预约外部核销
// taobao.life.reservation.trade.consume.notice
//
// 生服团购商品，购后预约。外部ISV进行核销
func TaobaoLifeReservationTradeConsumeNotice(clt *core.SDKClient, req *trade.TaobaoLifeReservationTradeConsumeNoticeAPIRequest, resp *trade.TaobaoLifeReservationTradeConsumeNoticeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
