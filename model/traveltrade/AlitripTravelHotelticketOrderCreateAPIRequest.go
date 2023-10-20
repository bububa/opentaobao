package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelhotelticketordercreateAPIRequest 创单(支付订单)通知 API请求
// alitrip.travel.hotelticket.order.create
//
// 创单(支付订单)通知
type AlitriptravelhotelticketordercreateAPIRequest struct {
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

// NewAlitriptravelhotelticketordercreateRequest 初始化AlitriptravelhotelticketordercreateAPIRequest对象
func NewAlitriptravelhotelticketordercreateRequest() *AlitriptravelhotelticketordercreateAPIRequest {
	return &AlitriptravelhotelticketordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelhotelticketordercreateAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelhotelticketordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelhotelticketordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendParams is ExtendParams Setter
// 扩展参数 支持的key: [&#34;hotelName&#34; : &#34;酒店名称&#34;, &#34;roomName&#34; : &#34;房型名称&#34;, &#34;productName&#34; : &#34;产品名称&#34;,  &#34;desc&#34; : &#34;描述&#34;] value字符长度不超过100
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetExtendParams(_extendParams string) error {
	r._extendParams = _extendParams
	r.Set("extend_params", _extendParams)
	return nil
}

// GetExtendParams ExtendParams Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetExtendParams() string {
	return r._extendParams
}

// SetOrderId is OrderId Setter
// 系统商订单号
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetFailMsg is FailMsg Setter
// 创单出票失败原因信息
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetFailMsg(_failMsg string) error {
	r._failMsg = _failMsg
	r.Set("fail_msg", _failMsg)
	return nil
}

// GetFailMsg FailMsg Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetFailMsg() string {
	return r._failMsg
}

// SetFliggyOrderId is FliggyOrderId Setter
// 飞猪订单号
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetFliggyOrderId(_fliggyOrderId string) error {
	r._fliggyOrderId = _fliggyOrderId
	r.Set("fliggy_order_id", _fliggyOrderId)
	return nil
}

// GetFliggyOrderId FliggyOrderId Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetFliggyOrderId() string {
	return r._fliggyOrderId
}

// SetVouchers is Vouchers Setter
// 凭证信息 无数据时传空集合
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetVouchers(_vouchers *HotelTicketVoucherDto) error {
	r._vouchers = _vouchers
	r.Set("vouchers", _vouchers)
	return nil
}

// GetVouchers Vouchers Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetVouchers() *HotelTicketVoucherDto {
	return r._vouchers
}

// SetStatus is Status Setter
// 创单结果状态 1：创单出票成功， 2：创单出票失败
func (r *AlitriptravelhotelticketordercreateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlitriptravelhotelticketordercreateAPIRequest) GetStatus() int64 {
	return r._status
}
