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
    expireTime   int64
    // 淘宝订单号
    tid   int64
    // 外部订单号
    outOrderId   string
    // 酒店code
    hotelCode   string
    // 酒店Id
    hid   int64
    // 请求报文
    context   string
    // 操作类型
    operateType   string
    // 请求唯一标识值
    requestId   string
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
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetExpireTime(expireTime int64) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

// ExpireTime Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetExpireTime() int64 {
    return r.expireTime
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetTid() int64 {
    return r.tid
}
// OutOrderId Setter
// 外部订单号
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetOutOrderId() string {
    return r.outOrderId
}
// HotelCode Setter
// 酒店code
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetHotelCode() string {
    return r.hotelCode
}
// Hid Setter
// 酒店Id
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetHid() int64 {
    return r.hid
}
// Context Setter
// 请求报文
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetContext() string {
    return r.context
}
// OperateType Setter
// 操作类型
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

// OperateType Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetOperateType() string {
    return r.operateType
}
// RequestId Setter
// 请求唯一标识值
func (r *TaobaoXhotelFutureSoftmodifyRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r TaobaoXhotelFutureSoftmodifyRequest) GetRequestId() string {
    return r.requestId
}
