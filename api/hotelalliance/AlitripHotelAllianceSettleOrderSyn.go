package hotelalliance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelAllianceSettleOrderSyn 菲住联盟分账成功订单同步
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
func AlitripHotelAllianceSettleOrderSyn(clt *core.SDKClient, req *hotelalliance.AlitripHotelAllianceSettleOrderSynAPIRequest, resp *hotelalliance.AlitripHotelAllianceSettleOrderSynAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
