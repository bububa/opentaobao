package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaolifereservationitemorderchange 生服购后预约单外部发起变更
// taobao.life.reservation.item.order.change
//
// 生服购后预约单外部发起变更，例如改期、取消。目前体检场景，用户会直接联系ISV改期/取消，因此开放给ISV这块的能力
func Taobaolifereservationitemorderchange(clt *core.SDKClient, req *trade.TaobaolifereservationitemorderchangeAPIRequest, session string) (*trade.TaobaolifereservationitemorderchangeAPIResponse, error) {
	var resp trade.TaobaolifereservationitemorderchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
