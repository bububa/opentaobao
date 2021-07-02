package hotelalliance

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelAllianceSettleOrderSynAPIRequest 菲住联盟分账成功订单同步 API请求
// alitrip.hotel.alliance.settle.order.syn
//
// 用于菲住联盟分账成功订单同步
type AlitripHotelAllianceSettleOrderSynAPIRequest struct {
	model.Params
	// 订单信息
	_orderInfo *AllianceSettleOrderInfo
}

// NewAlitripHotelAllianceSettleOrderSynRequest 初始化AlitripHotelAllianceSettleOrderSynAPIRequest对象
func NewAlitripHotelAllianceSettleOrderSynRequest() *AlitripHotelAllianceSettleOrderSynAPIRequest {
	return &AlitripHotelAllianceSettleOrderSynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.alliance.settle.order.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderInfo Setter
// 订单信息
func (r *AlitripHotelAllianceSettleOrderSynAPIRequest) SetOrderInfo(_orderInfo *AllianceSettleOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// Get OrderInfo Getter
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetOrderInfo() *AllianceSettleOrderInfo {
	return r._orderInfo
}
