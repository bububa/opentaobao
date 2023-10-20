package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketOrderCreateAPIRequest 创单(支付订单)通知 API请求
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
type AlitripTravelHotelticketOrderCreateAPIRequest struct {
	model.Params
	// 扩展参数 支持的key: ["hotelName" : "酒店名称", "roomName" : "房型名称", "productName" : "产品名称",  "desc" : "描述"] value字符长度不超过100
	_extendParams string
	// 系统商订单号
	_orderId string
	// 创单出票失败原因信息
	_failMsg string
	// 飞猪订单号
	_fliggyOrderId string
	// 凭证信息 无数据时传空集合
	_vouchers *HotelTicketVoucherDto
	// 创单结果状态 1：创单出票成功， 2：创单出票失败
	_status int64
}

// NewAlitripTravelHotelticketOrderCreateRequest 初始化AlitripTravelHotelticketOrderCreateAPIRequest对象
func NewAlitripTravelHotelticketOrderCreateRequest() *AlitripTravelHotelticketOrderCreateAPIRequest {
	return &AlitripTravelHotelticketOrderCreateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) Reset() {
	r._extendParams = ""
	r._orderId = ""
	r._failMsg = ""
	r._fliggyOrderId = ""
	r._vouchers = nil
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendParams is ExtendParams Setter
// 扩展参数 支持的key: [&#34;hotelName&#34; : &#34;酒店名称&#34;, &#34;roomName&#34; : &#34;房型名称&#34;, &#34;productName&#34; : &#34;产品名称&#34;,  &#34;desc&#34; : &#34;描述&#34;] value字符长度不超过100
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetOrderId is OrderId Setter
// 系统商订单号
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetFailMsg is FailMsg Setter
// 创单出票失败原因信息
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetFailMsg(_failMsg string) error {
	r._failMsg = _failMsg
	r.Set("fail_msg", _failMsg)
	return nil
}

// GetFailMsg FailMsg Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetFailMsg() string {
	return r._failMsg
}

// SetFliggyOrderId is FliggyOrderId Setter
// 飞猪订单号
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetFliggyOrderId(_fliggyOrderId string) error {
	r._fliggyOrderId = _fliggyOrderId
	r.Set("fliggy_order_id", _fliggyOrderId)
	return nil
}

// GetFliggyOrderId FliggyOrderId Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetFliggyOrderId() string {
	return r._fliggyOrderId
}

// SetVouchers is Vouchers Setter
// 凭证信息 无数据时传空集合
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetVouchers(_vouchers *HotelTicketVoucherDto) error {
	r._vouchers = _vouchers
	r.Set("vouchers", _vouchers)
	return nil
}

// GetVouchers Vouchers Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetVouchers() *HotelTicketVoucherDto {
	return r._vouchers
}

// SetStatus is Status Setter
// 创单结果状态 1：创单出票成功， 2：创单出票失败
func (r *AlitripTravelHotelticketOrderCreateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlitripTravelHotelticketOrderCreateAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlitripTravelHotelticketOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelHotelticketOrderCreateRequest()
	},
}

// GetAlitripTravelHotelticketOrderCreateRequest 从 sync.Pool 获取 AlitripTravelHotelticketOrderCreateAPIRequest
func GetAlitripTravelHotelticketOrderCreateAPIRequest() *AlitripTravelHotelticketOrderCreateAPIRequest {
	return poolAlitripTravelHotelticketOrderCreateAPIRequest.Get().(*AlitripTravelHotelticketOrderCreateAPIRequest)
}

// ReleaseAlitripTravelHotelticketOrderCreateAPIRequest 将 AlitripTravelHotelticketOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelHotelticketOrderCreateAPIRequest(v *AlitripTravelHotelticketOrderCreateAPIRequest) {
	v.Reset()
	poolAlitripTravelHotelticketOrderCreateAPIRequest.Put(v)
}
