package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
未来酒店信息下发 API请求
taobao.xhotel.future.softmodify

未来酒店信息下发，包含PMS订单查询和自助入住
*/
type TaobaoXhotelFutureSoftmodifyRequest struct {
    model.Params
    // 超时时长，默认3s
    _expireTime   int64
    // 淘宝订单号
    _tid   int64
    // 外部订单号
    _outOrderId   string
    // 酒店code
    _hotelCode   string
    // 酒店Id
    _hid   int64
    // 请求报文
    _context   string
    // 操作类型
    _operateType   string
    // 请求唯一标识值
    _requestId   string
}

// 初始化TaobaoXhotelFutureSoftmodifyRequest对象
func NewTaobaoXhotelFutureSoftmodifyRequest() *TaobaoXhotelFutureSoftmodifyRequest{
    return &TaobaoXhotelFutureSoftmodifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFutureSoftmodifyRequest) GetApiMethodName() string {
    return "taobao.xhotel.future.softmodify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFutureSoftmodifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExpireTime Setter
// 超时时长，默认3s
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetExpireTime(_expireTime int64) error {
    r._expireTime = _expireTime
    r.Set("expire_time", _expireTime)
    return nil
}

// ExpireTime Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetExpireTime() int64 {
    return r._expireTime
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetTid() int64 {
    return r._tid
}
// OutOrderId Setter
// 外部订单号
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetOutOrderId() string {
    return r._outOrderId
}
// HotelCode Setter
// 酒店code
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetHotelCode() string {
    return r._hotelCode
}
// Hid Setter
// 酒店Id
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetHid() int64 {
    return r._hid
}
// Context Setter
// 请求报文
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetContext(_context string) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetContext() string {
    return r._context
}
// OperateType Setter
// 操作类型
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOperateType(_operateType string) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetOperateType() string {
    return r._operateType
}
// RequestId Setter
// 请求唯一标识值
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetRequestId() string {
    return r._requestId
}
