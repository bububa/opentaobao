package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* TaobaoLifeReservationItemOrderChange
生服购后预约单外部发起变更
taobao.life.reservation.item.order.change

生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力 */
func TaobaoLifeReservationItemOrderChange(clt *core.SDKClient, req *trade.TaobaoLifeReservationItemOrderChangeAPIRequest, session string) (*trade.TaobaoLifeReservationItemOrderChangeAPIResponse, error) {
	var resp trade.TaobaoLifeReservationItemOrderChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
