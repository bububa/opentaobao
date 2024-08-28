package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoLifeReservationItemOrderConfirm 生服购后预约单外部确认
// taobao.life.reservation.item.order.confirm
//
// 生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回
func TaobaoLifeReservationItemOrderConfirm(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoLifeReservationItemOrderConfirmAPIRequest, resp *trade.TaobaoLifeReservationItemOrderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
