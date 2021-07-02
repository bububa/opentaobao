package hotelalliance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// AlitripHotelAllianceSettleOrderSyn 菲住联盟分账成功订单同步
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
func AlitripHotelAllianceSettleOrderSyn(clt *core.SDKClient, req *hotelalliance.AlitripHotelAllianceSettleOrderSynAPIRequest, session string) (*hotelalliance.AlitripHotelAllianceSettleOrderSynAPIResponse, error) {
	var resp hotelalliance.AlitripHotelAllianceSettleOrderSynAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
