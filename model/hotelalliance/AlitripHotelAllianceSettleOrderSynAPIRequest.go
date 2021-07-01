package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelAllianceSettleOrderSynAPIRequest
菲住联盟分账成功订单同步 API请求
alitrip.hotel.alliance.settle.order.syn

用于菲住联盟分账成功订单同步 */
type AlitripHotelAllianceSettleOrderSynAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *AllianceSettleOrderInfo
}

// New
