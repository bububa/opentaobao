package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelfuturesoftmodifyAPIRequest 未来酒店信息下发 API请求
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
type TaobaoxhotelfuturesoftmodifyAPIRequest struct {
	model.Params
	// 外部订单号
	_outOrderId string
	// 酒店code
	_hotelCode string
	// 请求报文
	_context string
	// 操作类型
	_operateType string
	// 请求唯一标识值
	_requestId string
	// 超时时长，默认3s
	_expireTime int64
	// 淘宝订单号
	_tid int64
	// 酒店Id
	_hid int64
}

// NewTaobaoxhotelfuturesoftmodifyRequest 初始化TaobaoxhotelfuturesoftmodifyAPIRequest对象
func NewTaobaoxhotelfuturesoftmodifyRequest() *TaobaoxhotelfuturesoftmodifyAPIRequest {
	return &TaobaoxhotelfuturesoftmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.future.softmodify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOrderId is OutOrderId Setter
// 外部订单号
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetHotelCode is HotelCode Setter
// 酒店code
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetContext is Context Setter
// 请求报文
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetContext(_context string) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetContext() string {
	return r._context
}

// SetOperateType is OperateType Setter
// 操作类型
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetRequestId is RequestId Setter
// 请求唯一标识值
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetExpireTime is ExpireTime Setter
// 超时时长，默认3s
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetExpireTime(_expireTime int64) error {
	r._expireTime = _expireTime
	r.Set("expire_time", _expireTime)
	return nil
}

// GetExpireTime ExpireTime Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetExpireTime() int64 {
	return r._expireTime
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetTid() int64 {
	return r._tid
}

// SetHid is Hid Setter
// 酒店Id
func (r *TaobaoxhotelfuturesoftmodifyAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelfuturesoftmodifyAPIRequest) GetHid() int64 {
	return r._hid
}
