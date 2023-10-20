package hotelalliance

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripHotelAllianceSettleOrderSynAPIRequest) Reset() {
	r._orderInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.alliance.settle.order.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 订单信息
func (r *AlitripHotelAllianceSettleOrderSynAPIRequest) SetOrderInfo(_orderInfo *AllianceSettleOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlitripHotelAllianceSettleOrderSynAPIRequest) GetOrderInfo() *AllianceSettleOrderInfo {
	return r._orderInfo
}

var poolAlitripHotelAllianceSettleOrderSynAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripHotelAllianceSettleOrderSynRequest()
	},
}

// GetAlitripHotelAllianceSettleOrderSynRequest 从 sync.Pool 获取 AlitripHotelAllianceSettleOrderSynAPIRequest
func GetAlitripHotelAllianceSettleOrderSynAPIRequest() *AlitripHotelAllianceSettleOrderSynAPIRequest {
	return poolAlitripHotelAllianceSettleOrderSynAPIRequest.Get().(*AlitripHotelAllianceSettleOrderSynAPIRequest)
}

// ReleaseAlitripHotelAllianceSettleOrderSynAPIRequest 将 AlitripHotelAllianceSettleOrderSynAPIRequest 放入 sync.Pool
func ReleaseAlitripHotelAllianceSettleOrderSynAPIRequest(v *AlitripHotelAllianceSettleOrderSynAPIRequest) {
	v.Reset()
	poolAlitripHotelAllianceSettleOrderSynAPIRequest.Put(v)
}
