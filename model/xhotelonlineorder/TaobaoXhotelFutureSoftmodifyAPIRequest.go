package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFutureSoftmodifyAPIRequest 未来酒店信息下发 API请求
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
type TaobaoXhotelFutureSoftmodifyAPIRequest struct {
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

// NewTaobaoXhotelFutureSoftmodifyRequest 初始化TaobaoXhotelFutureSoftmodifyAPIRequest对象
func NewTaobaoXhotelFutureSoftmodifyRequest() *TaobaoXhotelFutureSoftmodifyAPIRequest {
	return &TaobaoXhotelFutureSoftmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.future.softmodify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutOrderId is OutOrderId Setter
// 外部订单号
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetHotelCode is HotelCode Setter
// 酒店code
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetContext is Context Setter
// 请求报文
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetContext(_context string) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetContext() string {
	return r._context
}

// SetOperateType is OperateType Setter
// 操作类型
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetRequestId is RequestId Setter
// 请求唯一标识值
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetExpireTime is ExpireTime Setter
// 超时时长，默认3s
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetExpireTime(_expireTime int64) error {
	r._expireTime = _expireTime
	r.Set("expire_time", _expireTime)
	return nil
}

// GetExpireTime ExpireTime Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetExpireTime() int64 {
	return r._expireTime
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetTid() int64 {
	return r._tid
}

// SetHid is Hid Setter
// 酒店Id
func (r *TaobaoXhotelFutureSoftmodifyAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelFutureSoftmodifyAPIRequest) GetHid() int64 {
	return r._hid
}
