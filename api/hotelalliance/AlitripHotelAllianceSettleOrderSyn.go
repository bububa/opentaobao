package hotelalliance

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelAllianceSettleOrderSyn 菲住联盟分账成功订单同步
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
func AlitripHotelAllianceSettleOrderSyn(ctx context.Context, clt *core.SDKClient, req *hotelalliance.AlitripHotelAllianceSettleOrderSynAPIRequest, resp *hotelalliance.AlitripHotelAllianceSettleOrderSynAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
