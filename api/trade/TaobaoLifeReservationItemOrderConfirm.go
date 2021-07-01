package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* TaobaoLifeReservationItemOrderConfirm
生服购后预约单外部确认
taobao.life.reservation.item.order.confirm

生服团购下单预约后，用户改期/取消，外调ISV。ISV人工确认后调接口同意或驳回 */
func TaobaoLifeReservationItemOrderConfirm(clt *core.SDKClient, req *trade.TaobaoLifeReservationItemOrderConfirmAPIRequest, session string) (*trade.TaobaoLifeReservationItemOrderConfirmAPIResponse, error) {
	var resp trade.TaobaoLifeReservationItemOrderConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
