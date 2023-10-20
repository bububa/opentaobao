package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketOrderVerifyAPIRequest 订单核销通知 API请求
// alitrip.travel.hotelticket.order.verify
//
// 订单核销通知
type AlitripTravelHotelticketOrderVerifyAPIRequest struct {
	model.Params
	// 扩展参数
	_extendParams string
	// 系统商订单号
	_orderId string
	// 飞猪订单号
	_fliggyOrderId string
	// 凭证信息
	_vouchers *HotelTicketVerifyVoucherDto
}

// NewAlitripTravelHotelticketOrderVerifyRequest 初始化AlitripTravelHotelticketOrderVerifyAPIRequest对象
func NewAlitripTravelHotelticketOrderVerifyRequest() *AlitripTravelHotelticketOrderVerifyAPIRequest {
	return &AlitripTravelHotelticketOrderVerifyAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelHotelticketOrderVerifyAPIRequest) Reset() {
	r._extendParams = ""
	r._orderId = ""
	r._fliggyOrderId = ""
	r._vouchers = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendParams is ExtendParams Setter
// 扩展参数
func (r *AlitripTravelHotelticketOrderVerifyAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetOrderId is OrderId Setter
// 系统商订单号
func (r *AlitripTravelHotelticketOrderVerifyAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetFliggyOrderId is FliggyOrderId Setter
// 飞猪订单号
func (r *AlitripTravelHotelticketOrderVerifyAPIRequest) SetFliggyOrderId(_fliggyOrderId string) error {
	r._fliggyOrderId = _fliggyOrderId
	r.Set("fliggy_order_id", _fliggyOrderId)
	return nil
}

// GetFliggyOrderId FliggyOrderId Getter
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetFliggyOrderId() string {
	return r._fliggyOrderId
}

// SetVouchers is Vouchers Setter
// 凭证信息
func (r *AlitripTravelHotelticketOrderVerifyAPIRequest) SetVouchers(_vouchers *HotelTicketVerifyVoucherDto) error {
	r._vouchers = _vouchers
	r.Set("vouchers", _vouchers)
	return nil
}

// GetVouchers Vouchers Getter
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetVouchers() *HotelTicketVerifyVoucherDto {
	return r._vouchers
}

var poolAlitripTravelHotelticketOrderVerifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelHotelticketOrderVerifyRequest()
	},
}

// GetAlitripTravelHotelticketOrderVerifyRequest 从 sync.Pool 获取 AlitripTravelHotelticketOrderVerifyAPIRequest
func GetAlitripTravelHotelticketOrderVerifyAPIRequest() *AlitripTravelHotelticketOrderVerifyAPIRequest {
	return poolAlitripTravelHotelticketOrderVerifyAPIRequest.Get().(*AlitripTravelHotelticketOrderVerifyAPIRequest)
}

// ReleaseAlitripTravelHotelticketOrderVerifyAPIRequest 将 AlitripTravelHotelticketOrderVerifyAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelHotelticketOrderVerifyAPIRequest(v *AlitripTravelHotelticketOrderVerifyAPIRequest) {
	v.Reset()
	poolAlitripTravelHotelticketOrderVerifyAPIRequest.Put(v)
}
