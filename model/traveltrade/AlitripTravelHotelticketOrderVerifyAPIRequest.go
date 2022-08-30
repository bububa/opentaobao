package traveltrade

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelHotelticketOrderVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
